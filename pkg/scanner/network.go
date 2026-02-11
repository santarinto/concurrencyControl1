package scanner

import (
	"fmt"
	"net"
	"os/exec"
	"sync"
	"sync/atomic"
	"time"
)

type ScanResult struct {
	IP           string
	Hostname     string
	Ping         bool
	Port22       bool
	Port80       bool
	Port443      bool
	PortMySQL    bool
	PortRabbitMQ bool
	PortPostgres bool
}

type ProgressCallback func(current, total int, ip string)

func ScanNetwork(startIP, endIP string, progressCallback ProgressCallback) ([]ScanResult, error) {
	ips, err := generateIPRange(startIP, endIP)
	if err != nil {
		return nil, fmt.Errorf("failed to generate IP range: %w", err)
	}

	total := len(ips)
	var results []ScanResult
	var mu sync.Mutex
	var processed atomic.Int32

	semaphore := make(chan struct{}, 4)
	var wg sync.WaitGroup

	for _, ipStr := range ips {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(ipAddr string) {
			defer wg.Done()
			defer func() { <-semaphore }()

			result := scanHost(ipAddr)

			if result.Port22 || result.Port80 || result.Port443 {
				mu.Lock()
				results = append(results, result)
				mu.Unlock()
			}

			current := int(processed.Add(1))
			if progressCallback != nil {
				progressCallback(current, total, ipAddr)
			}
		}(ipStr)
	}

	wg.Wait()
	return results, nil
}

func generateIPRange(startIP, endIP string) ([]string, error) {
	start := net.ParseIP(startIP)
	end := net.ParseIP(endIP)

	if start == nil || end == nil {
		return nil, fmt.Errorf("invalid IP address")
	}

	var ips []string
	for ip := start; !ip.Equal(end); inc(ip) {
		ips = append(ips, ip.String())
	}
	ips = append(ips, end.String())

	return ips, nil
}

func scanHost(ip string) ScanResult {
	result := ScanResult{
		IP: ip,
	}

	result.Ping = pingHost(ip)

	ports := map[string]*bool{
		"22":  &result.Port22,
		"80":  &result.Port80,
		"443": &result.Port443,
	}

	for port, status := range ports {
		address := fmt.Sprintf("%s:%s", ip, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err == nil {
			conn.Close()
			*status = true
		}
	}

	if result.Port22 {
		dbPorts := map[string]*bool{
			"3306": &result.PortMySQL,
			"5672": &result.PortRabbitMQ,
			"5432": &result.PortPostgres,
		}

		for port, status := range dbPorts {
			address := fmt.Sprintf("%s:%s", ip, port)
			conn, err := net.DialTimeout("tcp", address, 1*time.Second)
			if err == nil {
				conn.Close()
				*status = true
			}
		}
	}

	if result.Port22 || result.Port80 || result.Port443 {
		result.Hostname = resolveHostname(ip)
	}

	return result
}

func pingHost(ip string) bool {
	cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
	err := cmd.Run()
	return err == nil
}

func resolveHostname(ip string) string {
	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		return ""
	}
	return names[0]
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
