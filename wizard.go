package logrus

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/thinkeridea/go-extend/exstrings"
)

const (
	c = "-config"
)

var (
	appName       = os.Args[0]
	ok            = false
	paths         = []string{" ", "", "etc", "opt", "data", "srv"}
	files         = []string{" ", "config", "dev_config", "test_config", "release_config"}
	fileType      = []string{"toml", "yaml"}
	kentRepoSlice = []string{
		"github.com/selead/",
		"github.com/vinerr/",
		"github.com/wpecker/",
		"github.com/oliveo/",
		"github.com/cuckoe/",
		"github.com/guavao/",
	}
	kentRepoMap = map[uint8]int{
		0: 18,
		1: 18,
		2: 19,
		3: 18,
		4: 18,
		5: 18,
	}
)

func init() {
	Infoln(os.Getenv("HOME"))
	Infoln(os.Args[0])
	Infoln(paths)

	if runtime.GOOS == "linux" {
		paths[1] = os.Getenv("HOME")[1:]
	}

	if strings.Contains(os.Args[0], `./`) {
		appName = exstrings.Replace(os.Args[0], `./`, "", -1)
	} else {
		if idx := strings.LastIndex(os.Args[0], "/"); idx != -1 {
			appName = exstrings.SubString(os.Args[0], idx+1, 0)
		}
	}
	Infoln(appName)

	pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err == nil {
		paths[0] = exstrings.SubString(pwd, 1, 0)
		files[0] = appName
	}

	customFormatter := new(TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	SetLevel(DebugLevel)
	SetFormatter(customFormatter)

	viper.SetConfigType("toml")
	if appName == "glide" {
		viper.AddConfigPath(os.Getenv("HOME"))
		addConfigPath()
		viper.SetConfigName("gwf")
		Infof("Using %s/gwf.toml for configuration", os.Getenv("HOME"))
		err := viper.ReadInConfig()
		ok = true
		if err != nil {
			Fatalf("Fatal error config file: %s", err)
		}
	} else {
		if len(os.Args) >= 2 {
			if v := strings.Split(os.Args[1], "="); len(v) == 2 && v[0] == c && exist(v[1]) {
				file := v[1]
				if strings.Contains(file, "/") {
					viper.AddConfigPath(path.Dir(file))
				} else {
					addConfigPath()
				}
				viper.SetConfigName(strings.Split(path.Base(file), ".")[0])
				err := viper.ReadInConfig()
				if err != nil {
					Fatalf("Fatal error config file: %s", err)
				}
				ok = true
				Infof("Using %s for configuration", file)
			}
		}
	}

	if !ok {
		addConfigPath()
		readConfig()
	}

	if viper.GetString("log.pathfile") != "" {
		setOutput(viper.GetString("log.pathfile"))
	}

	cfgWatch := viper.GetBool("config.watch")
	if !viper.IsSet("config.watch") {
		cfgWatch = true
	}

	if cfgWatch {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			Infoln("Config file changed:", e.Name)
		})
	}

	if viper.GetString("log.type") == "text" {
		customFormatter = new(TextFormatter)
		if viper.GetString("log.timestamp") != "" {
			customFormatter.TimestampFormat = viper.GetString("log.timestamp")
		}
		if !viper.IsSet("log.timestamp") {
			customFormatter.TimestampFormat = "01-02 15:04:05.000"
		}
		if viper.GetInt("log.msg_reserved") != 0 {
			customFormatter.MsgReservedWidth = viper.GetInt("log.msg_reserved")
		}
		if !viper.IsSet("log.msg_reserved") {
			customFormatter.MsgReservedWidth = 55
		}
		customFormatter.FullTimestamp = viper.GetBool("log.fulltime")
		if !viper.IsSet("log.fulltime") {
			customFormatter.FullTimestamp = true
		}
		SetFormatter(customFormatter)
	} else {
		customFormatter := new(JSONFormatter)
		if viper.GetString("log.timestamp") != "" {
			customFormatter.TimestampFormat = viper.GetString("log.timestamp")
		}
		if !viper.IsSet("log.timestamp") {
			customFormatter.TimestampFormat = "01-02 15:04:05.000"
		}
		SetFormatter(customFormatter)
	}

	Infoln("logrus & viper init done")

	hookFile := viper.GetBool("log.hookfile")
	if !viper.IsSet("log.hookfile") {
		hookFile = true
	}

	if hookFile {
		if level, err := ParseLevel(viper.GetString("log.level")); err == nil {
			SetLevel(level)
		} else {
			SetLevel(DebugLevel)
		}
		SetReportCaller(true)
		SetCallerPretty(defaultCallerPretty)
	} else {
		if level, err := ParseLevel(viper.GetString("log.level")); err == nil {
			SetLevel(level)
		} else {
			SetLevel(DebugLevel)
		}
	}
}

func defaultCallerPretty(frame *runtime.Frame) (function string, file string) {
	file = frame.File
	if idx := strings.LastIndex(file, "/"); idx != -1 {
		file = exstrings.SubString(file, idx+1, 0)
	}
	file = exstrings.Replace(file, ".go", "", -1)

	function = frame.Function
	if idx := strings.LastIndex(function, ".func"); idx != -1 {
		function = exstrings.SubString(function, 0, idx)
	}

	if idx := strings.Index(function, "github.com"); idx == -1 {
		if idx = strings.Index(function, "/"); idx != -1 {
			function = exstrings.SubString(function, idx+1, 0)
		}
	} else {
		for i, v := range kentRepoSlice {
			if idx := strings.Index(function, v); idx != -1 {
				Len := 18
				if i == 2 {
					Len = 19
				}
				function = exstrings.SubString(function, idx+Len, 0)
				if len(file) == 1 && exstrings.SubString(function, -2, 1) == "." {
					if idx := strings.Index(function, "/"); idx != -1 {
						function = exstrings.SubString(function, idx+1, 0)
					}
				} else {
					function = "^" + function
				}
				break
			}
		}
	}
	function = exstrings.Replace(function, ".(", ".", -1)
	function = exstrings.Replace(function, ").", ".", -1)
	function = exstrings.Replace(function, "main.", "m.", -1)
	file = fmt.Sprintf("%s:%d", file, frame.Line)
	return function, file
}

func SetCallerPretty(call func(*runtime.Frame) (function string, file string)) {
	FieldsLogger.Formatter.(*TextFormatter).CallerPrettyfier = call
}

func setOutput(file string) {
	if file == "APP.log" {
		file = exstrings.Replace(file, "APP", appName, -1)
	}

	if !viper.IsSet("log.max_size") &&
		!viper.IsSet("log.max_backups") &&
		!viper.IsSet("log.max_age") &&
		!viper.IsSet("log.compress") {
		f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			Fatalln(file, err)
		}
		SetOutput(f)
		Infoln("Set the output file location to:", file)
	} else {
		MaxSize := viper.GetInt("log.max_size")
		MaxBackups := viper.GetInt("log.max_backups")
		MaxAge := viper.GetInt("log.max_age")
		Compress := viper.GetBool("log.compress")

		if !viper.IsSet("log.max_size") {
			MaxSize = 80
		}
		if !viper.IsSet("log.max_backups") {
			MaxBackups = 15
		}
		if !viper.IsSet("log.max_age") {
			MaxAge = 30
		}
		if !viper.IsSet("log.compress") {
			Compress = true
		}

		SetOutput(&lumberjack.Logger{
			Filename:   file,
			MaxSize:    MaxSize, // megabytes
			MaxBackups: MaxBackups,
			MaxAge:     MaxAge,   // days
			Compress:   Compress, // disabled by default
		})
		f := "Start log rotation, the largest single file:%v, the largest backup:%v, save the number of days:%d，compress:%v.\n"
		Infof(f, MaxSize, MaxBackups, MaxAge, Compress)
		Infoln("Set the output file location to:", file)
	}

}

func addConfigPath() {
	for i, v := range paths {
		if i == 0 {
			viper.AddConfigPath(".")
		}
		viper.AddConfigPath(fmt.Sprintf("/%s/%s/", v, appName))
		viper.AddConfigPath(fmt.Sprintf("/%s/%s/etc/", v, appName))
	}
}

func readConfig() {
	var err error
	for _, t := range fileType {
		viper.SetConfigType(t)
		for _, f := range files {
			viper.SetConfigName(f)
			err = viper.ReadInConfig()
			file := ""
			if file = find(f, t); file != "" {
				if err != nil {
					Fatalln(file, err)
				}
			}
			if err == nil {
				Infof("No configuration file defined, Using %s for configuration", file)
				return
			}
		}
	}
	if err != nil {
		e := fmt.Sprintf("%s", err)
		e = exstrings.Replace(e, "release_config", exstrings.Join(files, ","), -1)
		Fatalf("%s", e)
	}
	return
}

func find(f, t string) string {
	file := fmt.Sprintf("%s.%s", f, t)
	if exist(file) {
		return file
	}
	for _, v := range paths {
		file := fmt.Sprintf("/%s/%s/%s.%s", v, appName, f, t)
		if exist(file) {
			return file
		}

		file = fmt.Sprintf("/%s/%s/etc/%s.%s", v, appName, f, t)
		if exist(file) {
			return file
		}
	}
	return ""
}

func exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
