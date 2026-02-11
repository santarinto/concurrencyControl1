package cmd

import (
	"concurrencyControl1/internal/config"
	"concurrencyControl1/pkg/display"
	"concurrencyControl1/pkg/scanner"
	"fmt"
	"os"
)

func RunScan(args []string) {
	cfg, err := config.LoadDefault()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Starting network scan...")
	fmt.Printf("IP Range: %s - %s\n", cfg.Network.StartIP, cfg.Network.EndIP)
	fmt.Printf("Subnet: %s\n\n", cfg.Network.Subnet)

	progressCallback := func(current, total int, ip string) {
		percent := float64(current) / float64(total) * 100
		fmt.Printf("\rScanning: %d/%d (%.1f%%) - Current IP: %s    ", current, total, percent, ip)
	}

	results, err := scanner.ScanNetwork(cfg.Network.StartIP, cfg.Network.EndIP, progressCallback)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError scanning network: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n\nScan completed!")

	if len(results) == 0 {
		fmt.Println("No hosts found")
		return
	}

	headers := []string{"No.", "IP Address", "Hostname", "Ping", "SSH", "HTTP", "HTTPS", "MySQL", "RabbitMQ", "Postgres"}
	var rows []display.TableRow

	for i, result := range results {
		rows = append(rows, display.TableRow{
			Columns: []string{
				fmt.Sprintf("%d", i+1),
				result.IP,
				result.Hostname,
				formatStatus(result.Ping),
				formatStatus(result.Port22),
				formatStatus(result.Port80),
				formatStatus(result.Port443),
				formatStatus(result.PortMySQL),
				formatStatus(result.PortRabbitMQ),
				formatStatus(result.PortPostgres),
			},
		})
	}

	fmt.Printf("\nFound %d host(s):\n\n", len(results))
	display.PrintTable(headers, rows)
}

func formatStatus(status bool) string {
	if status {
		return "*"
	}
	return ""
}
