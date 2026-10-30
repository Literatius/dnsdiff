package main

func CompareIP(oldIPs, newIPs []string) (added, removed, unchanged []string) {
	for _, oldip := range oldIPs {
		found := false
		for _, newip := range newIPs {
			if oldip == newip {
				found = true
				break
			}
		}
		if !found {
			removed = append(removed, oldip)
		} else {
			unchanged = append(unchanged, oldip)
		}
	}
	for _, newip := range newIPs {
		found := true
		for _, oldip := range oldIPs {
			if newip == oldip {
				found = false
				break
			}
		}
		if found {
			added = append(added, newip)
		}
	}

	return added, removed, unchanged
}
