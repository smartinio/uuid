package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/uuid"
)

var binary string

func TestMain(m *testing.M) {
	dir, _ := os.MkdirTemp("", "uuid-test")
	binary = filepath.Join(dir, "uuid")
	out, err := exec.Command("go", "build", "-o", binary, ".").CombinedOutput()
	if err != nil {
		panic("failed to build: " + string(out))
	}
	code := m.Run()
	os.RemoveAll(dir)
	os.Exit(code)
}

func TestV1(t *testing.T) {
	out, err := exec.Command(binary, "-v1").CombinedOutput()
	if err != nil {
		t.Fatalf("unexpected error: %s", out)
	}
	id, err := uuid.Parse(strings.TrimSpace(string(out)))
	if err != nil {
		t.Fatalf("invalid uuid: %q", out)
	}
	if id.Version() != 1 {
		t.Fatalf("expected version 1, got %d", id.Version())
	}
}

func TestV2(t *testing.T) {
	out, err := exec.Command(binary, "-v2").CombinedOutput()
	if err != nil {
		t.Fatalf("unexpected error: %s", out)
	}
	id, err := uuid.Parse(strings.TrimSpace(string(out)))
	if err != nil {
		t.Fatalf("invalid uuid: %q", out)
	}
	if id.Version() != 2 {
		t.Fatalf("expected version 2, got %d", id.Version())
	}
}

func TestV3(t *testing.T) {
	out, err := exec.Command(binary, "-v3", "example.com").CombinedOutput()
	if err != nil {
		t.Fatalf("unexpected error: %s", out)
	}
	id, err := uuid.Parse(strings.TrimSpace(string(out)))
	if err != nil {
		t.Fatalf("invalid uuid: %q", out)
	}
	if id.Version() != 3 {
		t.Fatalf("expected version 3, got %d", id.Version())
	}
}

func TestV3Deterministic(t *testing.T) {
	a, _ := exec.Command(binary, "-v3", "example.com").CombinedOutput()
	b, _ := exec.Command(binary, "-v3", "example.com").CombinedOutput()
	if strings.TrimSpace(string(a)) != strings.TrimSpace(string(b)) {
		t.Fatalf("expected deterministic output: %q != %q", a, b)
	}
}

func TestV3RequiresName(t *testing.T) {
	_, err := exec.Command(binary, "-v3").CombinedOutput()
	if err == nil {
		t.Fatal("expected error without name argument")
	}
}

func TestV4(t *testing.T) {
	out, err := exec.Command(binary, "-v4").CombinedOutput()
	if err != nil {
		t.Fatalf("unexpected error: %s", out)
	}
	id, err := uuid.Parse(strings.TrimSpace(string(out)))
	if err != nil {
		t.Fatalf("invalid uuid: %q", out)
	}
	if id.Version() != 4 {
		t.Fatalf("expected version 4, got %d", id.Version())
	}
}

func TestV5(t *testing.T) {
	out, err := exec.Command(binary, "-v5", "example.com").CombinedOutput()
	if err != nil {
		t.Fatalf("unexpected error: %s", out)
	}
	id, err := uuid.Parse(strings.TrimSpace(string(out)))
	if err != nil {
		t.Fatalf("invalid uuid: %q", out)
	}
	if id.Version() != 5 {
		t.Fatalf("expected version 5, got %d", id.Version())
	}
}

func TestV5Deterministic(t *testing.T) {
	a, _ := exec.Command(binary, "-v5", "example.com").CombinedOutput()
	b, _ := exec.Command(binary, "-v5", "example.com").CombinedOutput()
	if strings.TrimSpace(string(a)) != strings.TrimSpace(string(b)) {
		t.Fatalf("expected deterministic output: %q != %q", a, b)
	}
}

func TestV5RequiresName(t *testing.T) {
	_, err := exec.Command(binary, "-v5").CombinedOutput()
	if err == nil {
		t.Fatal("expected error without name argument")
	}
}

func TestV6(t *testing.T) {
	out, err := exec.Command(binary, "-v6").CombinedOutput()
	if err != nil {
		t.Fatalf("unexpected error: %s", out)
	}
	id, err := uuid.Parse(strings.TrimSpace(string(out)))
	if err != nil {
		t.Fatalf("invalid uuid: %q", out)
	}
	if id.Version() != 6 {
		t.Fatalf("expected version 6, got %d", id.Version())
	}
}

func TestV7(t *testing.T) {
	out, err := exec.Command(binary, "-v7").CombinedOutput()
	if err != nil {
		t.Fatalf("unexpected error: %s", out)
	}
	id, err := uuid.Parse(strings.TrimSpace(string(out)))
	if err != nil {
		t.Fatalf("invalid uuid: %q", out)
	}
	if id.Version() != 7 {
		t.Fatalf("expected version 7, got %d", id.Version())
	}
}
