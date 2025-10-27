package main

import (
  "fmt"
  "log"
  "os"
  "os/exec"
  "strings"

  "github.com/spf13/pflag"
)

func main(){
  installedPkgs := pflag.Bool("lp", false, "List third packages installed.")
  setPkg := pflag.String("pi", "", "Package Id to setup regysk frida config file")
  pflag.Parse()

  if *installedPkgs == true{
    listThirdPackages()
    os.Exit(0)
  }

  if *setPkg != "" {
    if checkConfig() {
    cleanConfig()
    }
    createNewConfig()
    setPackageId(*setPkg)
    os.Exit(0)
  }



}

func listThirdPackages(){
  cmd := exec.Command("adb", "shell", "pm", "list", "packages", "-3")
  out, err := cmd.Output()
  if err != nil {
    log.Fatalln(err)
  }

  splited := string(out)
  lines := strings.Split(splited, "\n")

  for _, line := range lines{
    cleanLine := strings.TrimPrefix(line, "package:")
    fmt.Println(cleanLine)
  }
}

func cleanConfig(){
  cmd := exec.Command("adb", "shell", "su -c rm /data/local/tmp/re.zyg.fri/config.json")
  err := cmd.Run()
  if err != nil {
    createNewConfig()
  }
}

func checkConfig() bool {
  cmd := exec.Command("adb", "shell", "su -c [ -e /data/local/tmp/re.zyg.fri/config.json ]")
  err := cmd.Run()
  if err != nil {
    return false
  }
  return true
}

func createNewConfig(){
  cmd := exec.Command("adb", "shell", "su -c cp /data/local/tmp/re.zyg.fri/config.json.example /data/local/tmp/re.zyg.fri/config.json")
  err := cmd.Run()
  if err != nil{
    log.Fatalln(err)
  }
}

func setPackageId(pkgName string){
  shellCmd := fmt.Sprintf("su -c sed -i s/com.example.package/%s/ /data/local/tmp/re.zyg.fri/config.json", pkgName)
  cmd := exec.Command("adb", "shell", shellCmd)
  err := cmd.Run()
  if err != nil{
    log.Fatalln(err)
  }

}
