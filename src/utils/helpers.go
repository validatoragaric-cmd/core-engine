package core_engine

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"

	"github.com/golang/protobuf/proto"
)

func GenerateRandomString() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func GetFileExt(filename string) string {
	ext := filepath.Ext(filename)
	return strings.ToLower(ext)
}

func IsNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func GetOsName() string {
	if runtime.GOOS == "windows" {
		return "Windows"
	}
	if runtime.GOOS == "darwin" { // macOS
		return "macOS"
	}
	return "Linux"
}

func GetOsArch() string {
	return runtime.GOARCH
}

func IsEmptySlice(s []string) bool {
	return len(s) == 0
}

func GetHash(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func StringSliceContains(a []string, s string) bool {
	for _, str := range a {
		if str == s {
			return true
		}
	}
	return false
}

func StringSliceIndex(a []string, s string) int {
	for i, str := range a {
		if str == s {
			return i
		}
	}
	return -1
}

func SortStrings(a []string) {
	sort.Strings(a)
}

func SortInts(a []int) {
	sort.Ints(a)
}

func GetMax(a []int) int {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

func GetMin(a []int) int {
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func GetUniqueValues(a []int) []int {
	sort.Ints(a)
	result := []int{a[0]}
	for i := 1; i < len(a); i++ {
		if a[i] != a[i-1] {
			result = append(result, a[i])
		}
	}
	return result
}

func GetCleanString(s string) string {
	var sb strings.Builder
	for _, r := range s {
		if unicode.IsSpace(r) {
			continue
		}
		sb.WriteRune(r)
	}
	return sb.String()
}

func IsEmptyString(s string) bool {
	return s == ""
}

func GetSlug(s string) string {
	slug := strings.ToLower(s)
	slug = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return r
		}
		return -1
	}, slug)
	return slug
}

func GetHexColor(r, g, b uint8) string {
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func GetRandomInt(min, max int64) int64 {
	return min + int64(math.Floor(rand.Float64()*float64(max-min+1)))
}

func GetRandomFloat(min, max float64) float64 {
	return min + math.Floor(rand.Float64()*float64(max-min))
}

func GetRandomFloat64() float64 {
	b := make([]byte, 8)
	rand.Read(b)
	return math.Float64frombits(uint64(binary.LittleEndian.Uint64(b)))
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func GetLocalHost() string {
	return os.Getenv("HOSTNAME")
}

func GetLocalPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func IsSameHost(req *http.Request, url string) bool {
	if req.URL.Host != url {
		return false
	}
	if req.URL.Scheme != "http" && req.URL.Scheme != "https" {
		return false
	}
	return true
}

func GetProtoMessageSize(message proto.Message) int {
	buf := new(bytes.Buffer)
	encoder := proto.NewEncoder(buf)
	err := encoder.Encode(message)
	if err != nil {
		log.Println(err)
		return 0
	}
	return buf.Len()
}

type once struct {
	mu sync.Mutex
	do func()
}

func (o *once) Do(f func()) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if o.do == nil {
		o.do = f
		f()
	}
}

func GetVersion() string {
	return "1.0"
}