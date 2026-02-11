package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

const usage = `Usage: uuid -v<version> [options]

Versions:
  -v1             Time-based (RFC 4122)
  -v2             DCE Security (uses current UID)
  -v3 <name>      Name-based MD5 (default namespace: dns)
  -v4             Random (RFC 4122)
  -v5 <name>      Name-based SHA1 (default namespace: dns)
  -v6             Reordered time-based (RFC 9562)
  -v7             Unix epoch time-based (RFC 9562)

Options for v3/v5:
  -ns <namespace>  Namespace: dns, url, oid, x500 (default: dns)`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	arg := os.Args[1]
	if arg == "-h" || arg == "--help" {
		fmt.Println(usage)
		return
	}

	if !strings.HasPrefix(arg, "-v") {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	versionStr := strings.TrimPrefix(arg, "-v")
	version, err := strconv.Atoi(versionStr)
	if err != nil || version < 1 || version > 7 {
		fatal("invalid version: %s (supported: 1-7)", versionStr)
	}

	remaining := os.Args[2:]

	var id uuid.UUID

	switch version {
	case 1:
		id, err = uuid.NewUUID()
	case 2:
		id, err = uuid.NewDCESecurity(uuid.Person, uint32(os.Getuid()))
	case 3:
		name, ns := parseNameArgs(remaining)
		id = uuid.NewMD5(ns, []byte(name))
	case 4:
		id = uuid.New()
	case 5:
		name, ns := parseNameArgs(remaining)
		id = uuid.NewSHA1(ns, []byte(name))
	case 6:
		id, err = uuid.NewV6()
	case 7:
		id, err = uuid.NewV7()
	}

	if err != nil {
		fatal("failed to generate UUID: %v", err)
	}

	fmt.Println(id)
}

func parseNameArgs(args []string) (string, uuid.UUID) {
	var name string
	ns := uuid.NameSpaceDNS

	for i := 0; i < len(args); i++ {
		if args[i] == "-ns" && i+1 < len(args) {
			ns = parseNamespace(args[i+1])
			i++
		} else if !strings.HasPrefix(args[i], "-") {
			name = args[i]
		}
	}

	if name == "" {
		fatal("v3 and v5 require a name argument\n\n%s", usage)
	}

	return name, ns
}

func parseNamespace(s string) uuid.UUID {
	switch strings.ToLower(s) {
	case "dns":
		return uuid.NameSpaceDNS
	case "url":
		return uuid.NameSpaceURL
	case "oid":
		return uuid.NameSpaceOID
	case "x500":
		return uuid.NameSpaceX500
	default:
		fatal("unknown namespace: %s (supported: dns, url, oid, x500)", s)
		return uuid.UUID{}
	}
}

func fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
