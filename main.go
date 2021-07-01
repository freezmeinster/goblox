package main

import (
    "fmt"
    "log"
    "github.com/go-ldap/ldap/v3"
)

func main(){

    ldapURL := "ldap://10.60.169.21:389"
    l, err := ldap.DialURL(ldapURL)
    if err != nil {
            log.Fatal(err)
    }
    defer l.Close()

    user := "fooUser"
    baseDN := "DC=example,DC=com"
    filter := fmt.Sprintf("(CN=%s)", ldap.EscapeFilter(user))

    // Filters must start and finish with ()!
    searchReq := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree, 0, 0, 0, false, filter, []string{"sAMAccountName"}, []ldap.Control{})

    result, err := l.Search(searchReq)
    if err != nil {
            log.Fatal(err)
    }

    log.Println("Got", len(result.Entries), "search results")
}
