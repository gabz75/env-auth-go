package core

import(
    "encoding/json"
    "io/ioutil"
    "os"
)

type Config struct {
    Db DBConfig
}

type DBConfig struct {
    Driver string
    Hostname string
    Database string
    Username string
    Password string
}

func GetConfig() Config {
    dat, err := ioutil.ReadFile("../config/" + env() + ".json")
    if err != nil {
        panic(err)
    }

    config := Config{}

    err = json.Unmarshal(dat, &config)

    if err != nil {
        panic(err)
    }

    return config
}


func env() string {
    env := os.Getenv("ENV")
    if env == "" {
        return "development"
    }
    return env
}
