package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindServers(t *testing.T) {
	sample := `ss://YWVzLTI1Ni1nY206WTZSOXBBdHZ4eHptR0M@54.36.174.181:5002
ss://YWVzLTI1Ni1nY206WTZSOXBBdHZ4eHptR0M@54.36.174.181:5001#Britain501%20%28t.me/Outline_Vpn%29
Other text
🇬🇧 #Britain

ss://YWVzLTI1Ni1nY206cEtFVzhKUEJ5VFZUTHRN@54.36.174.181:4444#Britain502%20%28t.me/Outline_Vpn%29
Definitely not clean
🇬🇧 #Britain

ss://YWVzLTI1Ni1nY206WTZSOXBBdHZ4eHptR0M@54.36.174.181:3306/Outline_Vpn%29

🇬🇧 #Britain

ss://YWVzLTI1Ni1nY206ZmFCQW9ENTRrODdVSkc3@54.36.174.181:2376#Britain504%20%28t.me

🇬🇧 #Britain

   ss://YWVzLTI1Ni1nY206UENubkg2U1FTbmZvUzI3QDUuMzkuNzAuMTM4OjgwOTA=#FrOutlineKeys

ss://YWVzLTI1Ni1nY206S2l4THZLendqZWtHMDBybQ==@ak1394.free.www.outline.network:8080#www.outline.network%20(japan)
vmess://eyJhZGQiOiIxMTYuMjAzLjczLjM0IiwiYWlkIjoiMCIsImFscG4iOiIiLCJob3N0IjoiZ29vZ2xlLmNvbSIsImlkIjoiMmFhYzc3ZGUtYjNlNC00MDE3LTg0NWMtY2ExMzgwZjJlOGQwIiwibmV0Ijoid3MiLCJwYXRoIjoiL3RlbGVncmFtLWlkLUBwcml2YXRldnBucyIsInBvcnQiOiI4MCIsInBzIjoiMTExNyhAT3V0bGluZV9WcG4pIiwic2N5IjoiYXV0byIsInNuaSI6IiIsInRscyI6IiIsInR5cGUiOiIiLCJ2IjoiMiJ9
vless://YWVzLTI1Ni1nY206ZmFCQW9ENTRrODdVSkc3@54.36.174.181:2376#Britain504%20%28t.me
`
	expectedServers := []string{
		"ss://YWVzLTI1Ni1nY206WTZSOXBBdHZ4eHptR0M@54.36.174.181:5002",
		"ss://YWVzLTI1Ni1nY206WTZSOXBBdHZ4eHptR0M@54.36.174.181:5001",
		"ss://YWVzLTI1Ni1nY206cEtFVzhKUEJ5VFZUTHRN@54.36.174.181:4444",
		"ss://YWVzLTI1Ni1nY206WTZSOXBBdHZ4eHptR0M@54.36.174.181:3306",
		"ss://YWVzLTI1Ni1nY206ZmFCQW9ENTRrODdVSkc3@54.36.174.181:2376",
		"ss://YWVzLTI1Ni1nY206UENubkg2U1FTbmZvUzI3QDUuMzkuNzAuMTM4OjgwOTA",
		"ss://YWVzLTI1Ni1nY206S2l4THZLendqZWtHMDBybQ==@ak1394.free.www.outline.network:8080",
	}

	servers := findServers(sample)
	assert.Len(t, servers, 7)
	assert.EqualValues(t, expectedServers, servers)
}
