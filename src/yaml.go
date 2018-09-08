package main

type Config struct {
    Application application `yaml:"application"`
}

type application struct {
    Db   dbinfo     `yaml:"db"`
    Ldap ldapinfo   `yaml:"ldap"`
}

type dbinfo struct {
    Url string     `yaml:"url"`
}

type ldapinfo struct {
    Host string     `yaml:"host"`
    Port string     `yaml:"port"`
    Dn   string     `yaml:"dn"`
    Password string `yaml:"password"`
}
