package main

import "testing"

func TestCompareIP(t *testing.T) {
	oldIPs := []string{"1.1.1.1", "8.8.8.8"}
	newIPs := []string{"1.1.1.1", "9.9.9.9"}

	added, removed, unchanged := CompareIP(oldIPs, newIPs)

	if len(added) != 1 {
		t.Errorf("Expected 1 added IP, got %d", len(added))
	} else if added[0] != "9.9.9.9" {
		t.Errorf("Expected 9.9.9.9, got %s", added[0])
	}

	if len(removed) != 1 {
		t.Errorf("Expected 1 removed IP, got %d", len(removed))
	} else if removed[0] != "8.8.8.8" {
		t.Errorf("Expected 8.8.8.8, got %s", removed[0])
	}

	if len(unchanged) != 1 {
		t.Errorf("Expected 1 unchanged IP, got %d", len(unchanged))
	} else if unchanged[0] != "1.1.1.1" {
		t.Errorf("Expected 1.1.1.1, got %s", unchanged[0])
	}
}

func TestCompareIPlen(t *testing.T) {
	oldIPs := []string{"1.1.1.1", "8.8.8.8"}
	newIPs := []string{"8.8.8.8", "1.1.1.1"}

	added, removed, unchanged := CompareIP(oldIPs, newIPs)

	if len(added) != 0 {
		t.Errorf("Expected 0 added IPs, got %d", len(added))
	}

	if len(removed) != 0 {
		t.Errorf("Expected 0 removed IPs, got %d", len(removed))
	}

	if len(unchanged) != 2 {
		t.Errorf("Expected 2 unchanged IPs, got %d", len(unchanged))
	}
}

func TestCompareIPEmptyOld(t *testing.T) {
	oldIPs := []string{}
	newIPs := []string{"1.1.1.1", "8.8.8.8"}

	added, removed, unchanged := CompareIP(oldIPs, newIPs)

	if len(added) != 2 {
		t.Errorf("Expected 2 added IPs, got %d", len(added))
	}

	if len(removed) != 0 {
		t.Errorf("Expected 0 removed IPs, got %d", len(removed))
	}

	if len(unchanged) != 0 {
		t.Errorf("Expected 0 unchanged IPs, got %d", len(unchanged))
	}
}

func TestCompareIPEmptyNew(t *testing.T) {
	oldIPs := []string{"1.1.1.1", "8.8.8.8"}
	newIPs := []string{}

	added, removed, unchanged := CompareIP(oldIPs, newIPs)

	if len(added) != 0 {
		t.Errorf("Expected 0 added IPs, got %d", len(added))
	}

	if len(removed) != 2 {
		t.Errorf("Expected 2 removed IPs, got %d", len(removed))
	}

	if len(unchanged) != 0 {
		t.Errorf("Expected 0 unchanged IPs, got %d", len(unchanged))
	}
}

func TestCompareIPBothEmpty(t *testing.T) {
	oldIPs := []string{}
	newIPs := []string{}

	added, removed, unchanged := CompareIP(oldIPs, newIPs)

	if len(added) != 0 || len(removed) != 0 || len(unchanged) != 0 {
		t.Errorf("Expected empty results, got added=%v, removed=%v, unchanged=%v",
			added, removed, unchanged)
	}
}
