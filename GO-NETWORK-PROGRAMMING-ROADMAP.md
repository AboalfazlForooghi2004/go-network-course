# نقشه راه جامع آموزش Go برای Network Programming

> **نسخه سند:** 1.0  
> **تاریخ تدوین:** 2026-11-20  
> **سیستم‌عامل آزمایشگاه:** Ubuntu Linux  
> **محیط توسعه:** Zsh  
> **نسخه Go محیط فعلی:** Go 1.27.4 / linux/amd64  
> **زمان پیشنهادی:** روزانه ۶۰ دقیقه، ۵ جلسه اصلی در هفته + یک جلسه اختیاری مرور  
> **مدت مسیر اصلی:** ۲۱ هفته، حدود ۱۰۵ ساعت آموزش و تمرین

---

## 1. هدف دوره

این دوره برای یادگیری Go با تمرکز بر **برنامه‌نویسی شبکه و ابزارسازی در Linux** طراحی شده است. هدف صرفاً یادگرفتن Syntax زبان یا نوشتن چند برنامه نمایشی نیست؛ در پایان مسیر باید بتوانیم برنامه‌هایی بسازیم که با Socketها، پروتکل‌های شبکه، Kernel لینوکس، Routing و ابزارهایی مانند Open vSwitch تعامل واقعی داشته باشند.

مسیر آموزشی از مبانی زبان شروع می‌شود و به‌تدریج به این حوزه‌ها می‌رسد:

- TCP، UDP، Unix Domain Socket و Socket Programming
- DNS، HTTP، HTTPS و TLS
- Concurrency، مدیریت Connection و طراحی سرویس‌های هم‌زمان
- ابزارهای عیب‌یابی و مانیتورینگ شبکه
- Linux Networking، Netlink، Interface، Address، Route، Neighbor و VRF
- Open vSwitch، OVSDB و مقدمات SDN
- Packet Capture و تحلیل بسته‌های شبکه
- تست، Logging، Metrics، Profiling، امنیت و آماده‌سازی برای Production

### خروجی نهایی مورد انتظار

ساخت و انتشار یک پروژه Go برای مشاهده و عیب‌یابی شبکه Linux که قابلیت‌هایی مانند نمایش Interfaceها، IPها و Routeها، بررسی DNS و TCP Connectivity، ارائه API و Metrics و در مراحل بعد مشاهده وضعیت OVS را داشته باشد.

---

## 2. پیش‌نیازها و محیط اجرا

### دانش و مهارت‌های لازم

- آشنایی مقدماتی با برنامه‌نویسی؛ تجربه قبلی حرفه‌ای با Go لازم نیست.
- آشنایی عملی مقدماتی با Linux و Terminal.
- آشنایی پایه با IP، Port و مفاهیم ابتدایی شبکه؛ مفاهیم عمیق‌تر در طول دوره مرور می‌شوند.

### محیط فعلی هنرجو

- Ubuntu Linux، معماری `x86_64`
- Zsh
- Go `1.27.4` در مسیر `/usr/local/go`
- مسیر اجرایی Go: `/usr/local/go/bin/go`

در ابتدای کار هر محیط یا ماشین آزمایشگاهی باید با دستورهای زیر بررسی شود:

```bash
go version
go env GOROOT GOPATH GOMOD
uname -m
lsb_release -a
```

**نکته نسخه:** نسخه Go نصب‌شده در زمان آغاز دوره 1.27.4 است. در صورت انتشار نسخه جدید یا اصلاحیه امنیتی، قبل از ارتقا باید Release Notes رسمی بررسی شود. مثال‌های دوره باید با نسخه‌ای که واقعاً در آزمایشگاه نصب است آزمایش شوند.

---

## 3. اصول آموزشی دوره

در این دوره هیچ کدی را بدون دانستن دلیل نوشتنش وارد پروژه نمی‌کنیم. آموزش هر مفهوم باید پاسخ پرسش‌های زیر را روشن کند:

| پرسش | توضیح مورد انتظار |
|---|---|
| What? | این دستور، نوع، تابع یا Package دقیقاً چیست؟ |
| Why? | چرا در این مسئله به آن نیاز داریم؟ |
| Input? | چه ورودی‌ها و پارامترهایی دریافت می‌کند؟ |
| Output? | چه خروجی‌هایی برمی‌گرداند یا چه اثری ایجاد می‌کند؟ |
| Error? | چه خطاهایی ممکن است رخ دهد و چگونه باید مدیریت شوند؟ |
| Runtime? | هنگام اجرا در Go، سیستم‌عامل یا Kernel چه اتفاقی می‌افتد؟ |
| Network? | این بخش در کدام لایه یا بخش از ارتباط شبکه نقش دارد؟ |
| Alternative? | چه راه‌های دیگری وجود دارد و چرا این راه انتخاب شده است؟ |

### قالب ثابت هر جلسه

1. **هدف یادگیری:** در پایان جلسه دقیقاً چه چیزی باید بلد باشیم؟
2. **مسئله واقعی:** این مفهوم کدام نیاز عملی شبکه را حل می‌کند؟
3. **توضیح مفهومی:** پیش از کدنویسی، مفهوم و جایگاه آن در معماری توضیح داده می‌شود.
4. **API و دستورها:** امضای تابع، پارامترها، خروجی، خطاها و محدودیت‌ها بررسی می‌شود.
5. **پیاده‌سازی مرحله‌ای:** برنامه به بخش‌های کوچک تقسیم می‌شود؛ از ارائه یک فایل کامل بدون توضیح پرهیز می‌کنیم.
6. **اجرای عملی:** دستورهای اجرا در Ubuntu و روش مشاهده رفتار برنامه ارائه می‌شود.
7. **آزمایش و خطایابی:** حالت موفق، ورودی نامعتبر، قطع ارتباط، Timeout و خطاهای محتمل بررسی می‌شوند.
8. **تمرین مستقل:** پاسخ تمرین از ابتدا داده نمی‌شود؛ ابتدا هنرجو تلاش می‌کند و سپس کدش بررسی می‌شود.
9. **خروجی و معیار عبور:** چک‌لیست پایان جلسه و شرط ورود به مبحث بعدی مشخص می‌شود.
10. **Git:** یک پیام Commit مناسب برای ثبت تغییرات پیشنهاد می‌شود.

### قواعد طراحی مثال‌ها

- ابتدا از کتابخانه استاندارد Go استفاده می‌کنیم.
- Dependency خارجی فقط وقتی اضافه می‌شود که نیاز فنی مشخص و دلیل انتخاب آن توضیح داده شده باشد.
- کدها باید قابل اجرا، قالب‌بندی‌شده با `gofmt` و تا حد مناسب تست‌پذیر باشند.
- در برنامه شبکه، خطا، Timeout، Cancellation، بستن منابع و رفتار در زمان قطع ارتباط بخشی از طراحی‌اند؛ موارد فرعی محسوب نمی‌شوند.
- عملیات تغییردهنده شبکه در محیط آزمایشگاهی و Network Namespace انجام می‌شود؛ از تغییر Route یا تنظیمات اصلی سیستم بدون طرح بازگشت و تأیید خودداری می‌کنیم.

---

## 4. نقشه راه ۲۱ هفته‌ای

### مرحله اول — تسلط بر مبانی Go

**مدت:** هفته‌های ۱ تا ۴  
**هدف:** ساختن پایه کافی برای نوشتن برنامه Go ساختاریافته، خوانا، قابل تست و قابل خطایابی.

#### هفته ۱ — Syntax و ساختار اولیه Go

| جلسه | سرفصل‌ها | دلیل و کاربرد |
|---|---|---|
| 1 | Go Toolchain، `go run`، `go build`، کامپایلر و Runtime | شناخت چرخه ساخت و اجرای برنامه و تفاوت اجرای مستقیم با ساخت Binary |
| 2 | `package`، `import`، تابع `main` | درک نقطه ورود برنامه و سازمان‌دهی کد |
| 3 | متغیر، ثابت، `var`، `:=`، Scope و Type Inference | نگهداری داده و فهم محدوده اعتبار متغیر |
| 4 | انواع داده `string`، `bool`، اعداد صحیح و اعشاری، تبدیل نوع | انتخاب نوع مناسب برای IP متنی، Port عددی، وضعیت Link و مقادیر اندازه‌گیری |
| 5 | `if`، `switch`، `for`، `range` و عملگرها | پیاده‌سازی تصمیم‌گیری و تکرار در منطق برنامه |

**مباحث تکمیلی هفته:** تفاوت `fmt.Print`، `fmt.Println` و `fmt.Printf`، قالب‌های `%s` و `%d`، کاراکتر newline، خطاهای حروف بزرگ و کوچک در نام‌ها، `gofmt` و خواندن خطای Compiler.

**تمرین/خروجی:** برنامه CLI ساده‌ای که Host، Port، Protocol و نام سرویس را تعریف و نمایش دهد. هدف این تمرین یادگیری متغیرها و ساختار برنامه است؛ هنوز هیچ ارتباط شبکه‌ای برقرار نمی‌شود.

#### هفته ۲ — ساختار داده‌ها و مدل‌سازی اطلاعات شبکه

- Array و تفاوت آن با Slice
- ساخت Slice، برش، `len`، `cap`، `append` و کپی داده
- Map، کلید و مقدار، بررسی وجود کلید
- Struct و تعریف مدل‌های دامنه
- Pointer، آدرس حافظه، Dereference و مقدار `nil`
- Method و Receiver
- String، Byte و Rune
- مفهوم Mutability و کپی/اشتراک داده در حد لازم

**چرا؟** برنامه مانیتورینگ باید بتواند تعداد زیادی Interface، IP، Route یا Connection را به‌صورت مجموعه‌ای از داده‌های ساختاریافته نگهداری کند. Struct مدل هر موجودیت را تعریف می‌کند و Slice/Map مجموعه آن‌ها را نگه می‌دارند.

**تمرین/خروجی:** تعریف مدل‌های `NetworkInterface`، `IPAddress` و `Route` و نگهداری چند نمونه از آن‌ها در Slice و Map.

#### هفته ۳ — توابع، خطاها، Interface و I/O

- تعریف تابع، پارامتر و چند مقدار بازگشتی
- تفاوت Return Value و Side Effect
- الگوی رایج بازگرداندن `value, error`
- مدیریت خطا با `if err != nil`
- ساخت خطا با `errors.New` و `fmt.Errorf`
- Wrap کردن خطا با `%w`
- بررسی خطا با `errors.Is` و `errors.As`
- Interface و پیاده‌سازی ضمنی آن در Go
- Pointer Receiver و Value Receiver
- `defer` و آزادسازی منابع
- `io.Reader` و `io.Writer`
- `bufio.Reader`، `bufio.Writer`، `bytes.Buffer`

**چرا؟** عملیات شبکه دائماً با خطا مواجه می‌شود: مقصد در دسترس نیست، پورت بسته است، Deadline پایان می‌یابد یا طرف مقابل اتصال را می‌بندد. علاوه بر این، بسیاری از APIهای Go داده را از طریق Reader و Writer جابه‌جا می‌کنند.

**تمرین/خروجی:** طراحی توابع اعتبارسنجی Host و Port و نوشتن تابعی که خطا را با زمینه کافی به لایه فراخواننده برگرداند.

#### هفته ۴ — Module، Package، CLI و تست

- مفهوم Module و فایل `go.mod`
- Package، Import Path، Exported و Unexported Identifier
- ساختار پوشه و مرزبندی مسئولیت‌ها
- آرگومان خط فرمان با `flag`
- خواندن تنظیمات و اعتبارسنجی ورودی
- Unit Test با Package `testing`
- Table-Driven Tests
- Test Helper و تست توابع خطاپذیر
- تست پوشش کد با `go test -cover`
- `go test`، `go vet`، `go build` و `go doc`
- آشنایی اولیه با Mock و Dependency Injection

**چرا؟** کدی که فقط روی سیستم نویسنده اجرا شود کافی نیست. باید بتوان آن را در Packageهای مجزا نگهداری، تست و توسعه داد.

**پروژه مرحله:** ابزار خط فرمان برای دریافت، اعتبارسنجی و نمایش مشخصات مقصد شبکه.

---

### مرحله دوم — Concurrency و مبانی شبکه

**مدت:** هفته‌های ۵ و ۶  
**هدف:** فهم اجرای هم‌زمان در Go و مرور مفاهیم شبکه‌ای که برای Socket Programming لازم‌اند.

#### هفته ۵ — Concurrency در Go

- Process، Thread و Goroutine؛ تفاوت مفهومی آن‌ها
- اجرای Goroutine با `go`
- Channelهای Unbuffered و Buffered
- ارسال و دریافت از Channel
- بستن Channel و الگوی `range`
- دستور `select`
- `sync.WaitGroup`
- `sync.Mutex` و `sync.RWMutex`
- Race Condition و اجرای `go test -race`
- Context، Cancellation، Deadline و Timeout
- جلوگیری از Goroutine Leak
- Worker Pool و محدودسازی هم‌زمانی
- مدیریت خطا در کارهای موازی

**چرا؟** سرور شبکه باید بتواند به چندین اتصال یا درخواست رسیدگی کند. اجرای بی‌قاعده Goroutine به‌تنهایی راه‌حل نیست؛ باید تعداد کارها، عمر آن‌ها، خطاها و لغو عملیات مدیریت شوند.

**پروژه:** Worker Pool برای اجرای هم‌زمان چند بررسی Connectivity روی مقصدهای آزمایشی، با محدودیت تعداد Worker و امکان Cancel.

#### هفته ۶ — مبانی شبکه برای برنامه‌نویسان

- مدل OSI و TCP/IP و نقش هر لایه
- Encapsulation و Decapsulation
- Ethernet Frame، MAC Address و MTU
- IPv4، Subnet Mask، CIDR و Default Gateway
- IPv6، Prefix و آدرس‌های اصلی
- ARP در IPv4 و Neighbor Discovery در IPv6
- Routing Table، Next Hop و Longest Prefix Match
- TCP، UDP و ICMP/ICMPv6
- Port، Socket Address و مفهوم اتصال
- DNS، Resolver و Resource Record
- NAT و Firewall در حد معماری
- تفاوت Control Plane و Data Plane در سطح مقدماتی

**ابزارهای آزمایشگاهی Linux:**

```bash
ip addr
ip route
ip neigh
ss -lntup
ping
traceroute
dig
curl
nc
tcpdump
```

**چرا؟** وقتی برنامه Go به مقصد متصل می‌شود، سیستم‌عامل باید Route انتخاب کند، همسایه لایه دوم را پیدا کند و بسته را ارسال کند. شناخت این مسیر برای عیب‌یابی برنامه شبکه ضروری است.

**خروجی:** مستندسازی مسیر یک ارتباط ساده از Application تا Interface خروجی با کمک ابزارهای Linux.

---

### مرحله سوم — Socket Programming

**مدت:** هفته‌های ۷ تا ۱۰  
**هدف:** ساخت ارتباط واقعی بین دو برنامه و فهم رفتار TCP، UDP و Unix Socket.

#### هفته ۷ — TCP Server و TCP Client

- مفهوم Socket و Socket Address
- نقش TCP و ارتباط Connection-Oriented
- TCP Three-Way Handshake
- تفاوت Bind، Listen، Accept و Connect
- Package استاندارد `net`
- `net.Listen`
- متد `Accept`
- `net.Dial` و `net.DialTimeout`
- `net.Conn` و متدهای `Read`، `Write` و `Close`
- تشخیص خطاهای اتصال و اشغال بودن Port
- مشاهده Listener با `ss`
- مشاهده Handshake و بسته‌ها با `tcpdump`

**پروژه:** TCP Echo Server و TCP Client. کلاینت پیام می‌فرستد و سرور همان داده را برمی‌گرداند.

#### هفته ۸ — جریان داده و Framing در TCP

- TCP یک Byte Stream است، نه Message Queue
- Partial Read و Partial Write
- مفهوم EOF
- مدیریت داده ناقص و چند پیام پشت سر هم
- `io.Reader`، `io.Writer` و `io.ReadFull`
- `bufio.Reader` و `bufio.Writer`
- Framing بر اساس Delimiter
- Framing بر اساس Length Prefix
- تبدیل عدد و بایت با `encoding/binary`
- Read Deadline و Write Deadline
- تشخیص قطع ارتباط و مدیریت منابع

**چرا؟** یک فراخوانی `Read` الزاماً یک پیام کامل برنامه را تحویل نمی‌دهد. اگر Framing را درست طراحی نکنیم، پیام‌ها ممکن است تکه‌تکه یا به‌صورت ترکیبی دریافت شوند.

**پروژه:** ارتقای Echo Server به پروتکلی با قالب پیام مشخص، محدودیت طول پیام و مدیریت پیام ناقص.

#### هفته ۹ — TCP Server هم‌زمان

- مدیریت چند Client
- Goroutine به‌ازای Connection و محدودیت‌های آن
- Connection Lifecycle
- مدیریت قطع ناگهانی Client
- محدودیت تعداد اتصال‌ها
- Timeout و Idle Timeout
- Context و Cancellation
- Graceful Shutdown
- Signal Handling
- جلوگیری از Goroutine و Resource Leak
- تست اتصال‌های موازی

**پروژه:** TCP Chat Server با چند Client و مدیریت Join/Leave.

#### هفته ۱۰ — UDP و Unix Domain Socket

**UDP:**

- Datagram و حفظ مرز Packet
- تفاوت UDP با TCP
- `net.ListenUDP`
- `ReadFromUDP` و `WriteToUDP`
- `net.PacketConn`
- Packet Loss، Reordering و Duplicate
- محدودیت اندازه Datagram
- Sequence Number و Acknowledgment در یک پروتکل آموزشی
- Timeout و Retry

**Unix Domain Socket:**

- ارتباط بین Processهای محلی
- `net.Listen("unix", ...)`
- مجوز فایل Socket
- کاربرد در سرویس‌های Linux

**پروژه:** UDP Echo Service و یک سرویس محلی مبتنی بر Unix Domain Socket.

---

### مرحله چهارم — پروتکل‌های کاربردی

**مدت:** هفته‌های ۱۱ تا ۱۳  
**هدف:** استفاده از Go برای پیاده‌سازی و بررسی پروتکل‌هایی که در سرویس‌ها و شبکه‌های واقعی استفاده می‌شوند.

#### هفته ۱۱ — DNS و IP Address Programming

- ساختار DNS Query و Response
- Resolver و Name Resolution
- Recursive و Iterative Query
- A و AAAA
- CNAME و MX
- NS و TXT
- PTR و Reverse Lookup
- TTL و Cache
- DNS روی UDP و TCP
- `net.Resolver`
- `net.LookupHost` و `net.LookupIP`
- Package `net/netip`
- `netip.ParseAddr`، `netip.ParsePrefix` و `netip.AddrPort`
- اعتبارسنجی IPv4، IPv6 و CIDR

**پروژه:** DNS Inspector برای بررسی نام دامنه، آدرس‌ها و Resource Recordهای موردنیاز و نمایش خروجی ساختاریافته.

#### هفته ۱۲ — HTTP Client و HTTP Server

- ساختار Request و Response
- Method، Status Code، Header و Body
- URL و Query Parameter
- HTTP/1.1 و آشنایی با HTTP/2
- JSON Encoding و Decoding
- `net/http` و `http.Handler`
- `http.HandlerFunc`
- Router و Middleware
- ساخت HTTP Client
- مدیریت و بستن Response Body
- Client Timeout
- Request Context
- Server Read/Write/Idle Timeout
- تست با `httptest`

**پروژه:** HTTP API برای نمایش وضعیت شبکه و اجرای Checkهای ساده Connectivity.

#### هفته ۱۳ — TLS و ارتباط امن

- هدف TLS و مدل Secure Channel
- رمزنگاری متقارن و نامتقارن در حد مفهومی
- Public Key و Private Key
- Certificate، CA و Certificate Chain
- Server Authentication و Client Authentication
- TLS Handshake
- TLS 1.3
- Mutual TLS یا mTLS
- Certificate Verification و خطاهای رایج
- `crypto/tls`
- HTTPS با `net/http`
- مدیریت امن فایل‌های Certificate و Key

**پروژه:** HTTPS Client/Server محلی و سپس برقراری mTLS با CA آزمایشگاهی.

---

### مرحله پنجم — ابزارهای کاربردی شبکه

**مدت:** هفته‌های ۱۴ و ۱۵  
**هدف:** تبدیل دانش پروتکل و هم‌زمانی به ابزارهای تشخیصی قابل استفاده.

#### هفته ۱۴ — Port Scanner

- هدف و محدودیت Port Scanning
- TCP Connect Scan
- تفاوت Open، Closed، Timeout و احتمال Filtered بودن
- `net.Dialer`
- Timeout و Deadline
- Concurrent Scan
- Worker Pool و Rate Limiting
- Cancellation
- مدیریت Socket و محدودیت منابع
- خروجی Text و JSON
- تست دقت و رفتار روی شبکه کنترل‌شده

**پروژه:** Port Scanner با Host، محدوده Port، Timeout و تعداد Worker قابل تنظیم.

**محدوده مجاز:** اسکن فقط روی سیستم‌ها و شبکه‌هایی که مالک آن‌ها هستیم یا مجوز صریح آزمایش آن‌ها را داریم.

#### هفته ۱۵ — Network Monitor

- فهرست Interfaceهای شبکه
- IP Address و MAC Address
- وضعیت Up/Down Interface
- Gateway و Route
- DNS Health Check
- TCP Connectivity Check
- اندازه‌گیری Latency
- اجرای دوره‌ای Probeها
- مدل‌سازی نتیجه و وضعیت مقصد
- Context و Shutdown
- Logging و JSON Output
- آشنایی با Metrics

**پروژه:** Network Monitor برای بررسی دوره‌ای چند مقصد، DNS و وضعیت Interfaceها.

---

### مرحله ششم — Linux Network Programming

**مدت:** هفته‌های ۱۶ و ۱۷  
**هدف:** عبور از APIهای سطح Application و تعامل با زیرساخت شبکه Kernel لینوکس.

#### هفته ۱۶ — Linux Netlink

- معماری Linux Networking
- Kernel Space و User Space
- Netlink Socket و خانواده‌های Netlink
- Routing Netlink یا RTNL
- Interface، Address، Route و Neighbor Management
- Network Namespace
- Linux Capabilities و تفاوت آن با اجرای کامل به‌عنوان Root
- ساختار پیام Netlink در حد لازم
- `golang.org/x/sys/unix`
- کتابخانه `github.com/vishvananda/netlink`
- خواندن اطلاعات Kernel و مدیریت خطا/Permission

**پروژه:** ابزار خواندنی برای نمایش Interface، Address، Neighbor و Route از طریق Netlink.

#### هفته ۱۷ — Routing، Policy Routing و VRF

- Linux Routing Table
- Route، Gateway و Next Hop
- Route Metric
- Longest Prefix Match
- Policy-Based Routing
- `ip rule` و چند Routing Table
- ECMP در سطح مفهومی و عملی
- VRF در Linux
- ارتباط VRF با Routing Table
- Network Namespace و جداسازی شبکه
- IPv4 و IPv6 Route
- مشاهده و تغییر Route از طریق Netlink
- Rollback و مدیریت خطا در عملیات تغییردهنده

**پروژه:** Route Inspector و سپس Route Manager آزمایشی در Network Namespace اختصاصی.

**محدودیت ایمنی:** در طول یادگیری، Routeهای سیستم اصلی تغییر داده نمی‌شوند. عملیات تغییردهنده ابتدا در Namespace آزمایشگاهی و با روش بازگشت آزمایش می‌شود.

---

### مرحله هفتم — OVS، SDN، Packet Analysis و Production

**مدت:** هفته‌های ۱۸ تا ۲۰  
**هدف:** مطالعه سیستم‌های شبکه‌ای واقعی، کنترل OVS، تحلیل Packet و آماده‌سازی برنامه برای استفاده پایدار.

#### هفته ۱۸ — Open vSwitch و OVSDB

- معماری Open vSwitch
- Bridge، Port و Interface
- Datapath و Flow Table
- Control Plane و Data Plane
- تفاوت OVSDB و OpenFlow
- JSON-RPC
- OVSDB Schema
- Monitor و Transaction
- مدل‌سازی داده OVSDB
- کتابخانه `github.com/ovn-kubernetes/libovsdb`
- اتصال به OVSDB Server
- خواندن وضعیت Bridge، Port و Interface
- تغییر Configuration در Lab کنترل‌شده

**پروژه:** OVSDB Inspector برای اتصال به OVSDB و نمایش Bridge، Port و Interface.

#### هفته ۱۹ — Packet Capture و تحلیل ترافیک

- Packet Capture و تفاوت آن با Decode
- Promiscuous Mode
- PCAP و فایل Capture
- Ethernet، IPv4/IPv6، TCP، UDP و ICMP Header
- Packet Filtering
- BPF Filter در حد استفاده عملی
- تحلیل Flow و Conversation
- کتابخانه `github.com/google/gopacket`
- محدودیت Permission در Packet Capture لینوکس
- تطبیق خروجی برنامه با `tcpdump` و Wireshark

**پروژه:** Packet Analyzer برای خواندن فایل PCAP و استخراج اطلاعات Flowها و پروتکل‌ها.

#### هفته ۲۰ — Production Readiness و Observability

- Structured Logging با `log/slog`
- Metrics و Prometheus
- Health Check و Readiness
- Retry و Exponential Backoff
- Rate Limiting
- Graceful Shutdown
- Context Propagation
- مدیریت منابع و جلوگیری از Leak
- Unit، Integration و Fuzz Test
- Benchmark با `testing.B`
- Profiling با `pprof`
- Race Detection
- Dependency Management
- بررسی آسیب‌پذیری Dependencyها
- CI با GitHub Actions
- مستندسازی، Configuration و Release

**خروجی:** یکی از پروژه‌های قبلی را به سطحی برسانیم که تست، Log، Metrics، Shutdown و راهنمای اجرای قابل اتکا داشته باشد.

---

### مرحله هشتم — پروژه نهایی

**مدت:** هفته ۲۱  
**عنوان:** Go Network Diagnostics & Management Agent

#### قابلیت‌های هدف

| قابلیت | فناوری پیشنهادی |
|---|---|
| نمایش Interfaceها | Linux Netlink |
| نمایش IP و CIDR | `net/netip` |
| نمایش Routing Table | Netlink |
| بررسی TCP Connectivity | `net.Dialer` |
| بررسی DNS | `net.Resolver` |
| بررسی HTTP/HTTPS | `net/http` |
| اجرای مانیتورینگ دوره‌ای | Goroutine و Context |
| ارائه API | `net/http` |
| ارائه Metrics | Prometheus |
| گزارش‌گیری | `log/slog` |
| مشاهده وضعیت OVS | OVSDB و `libovsdb` |

#### فازبندی پروژه نهایی

1. تعریف نیازمندی‌ها و مدل داده
2. طراحی Packageها و مرز مسئولیت‌ها
3. پیاده‌سازی قابلیت‌های Read-only شبکه
4. افزودن Probeهای DNS و TCP
5. افزودن API و Health Check
6. افزودن Logging و Metrics
7. تست Unit و Integration
8. آزمایش در Ubuntu و Network Namespace
9. مستندسازی معماری و روش اجرا
10. ساخت Release و انتشار مخزن

**اصل ایمنی:** نسخه اول پروژه فقط خواندنی است. افزودن عملیات مدیریتی مانند تغییر Route یا OVS Configuration به طراحی مجوز، اعتبارسنجی، Log، تست و روش Rollback نیاز دارد و فقط در Lab انجام می‌شود.

---

## 5. خروجی‌های پروژه‌ای دوره

| شماره | پروژه | خروجی قابل ارائه |
|---|---|---|
| 1 | Network Information CLI | تمرین متغیر، Struct، ورودی و خروجی |
| 2 | Concurrent Connectivity Worker | تمرین Goroutine، Channel، Context و Timeout |
| 3 | TCP Echo Server/Client و Chat Server | Socket Programming و مدیریت Connection |
| 4 | DNS Inspector و HTTP API | پروتکل‌های Application و Resolver |
| 5 | Port Scanner و Network Monitor | ابزارهای تشخیصی شبکه |
| 6 | Linux Route Inspector و OVSDB Inspector | Kernel Networking، Netlink و OVS |
| 7 | Packet Analyzer و پروژه نهایی | تحلیل ترافیک و یکپارچه‌سازی مهارت‌ها |

---

## 6. ساختار پیشنهادی Repository

نام پیشنهادی مخزن: `go-networking-lab`

ساختار پیشنهادی برای زمانی که چند درس و پروژه شکل گرفت:

```text
go-networking-lab/
├── README.md
├── ROADMAP.md
├── LICENSE
├── go.mod
├── go.sum
├── .gitignore
├── docs/
│   ├── architecture/
│   ├── diagrams/
│   ├── network-fundamentals/
│   └── references.md
├── lessons/
│   ├── 01-go-basics/
│   ├── 02-data-types/
│   ├── 03-functions-errors/
│   ├── 04-packages-testing/
│   ├── 05-concurrency/
│   └── ...
├── cmd/
│   ├── netinfo/
│   ├── tcp-echo-server/
│   ├── tcp-echo-client/
│   ├── dns-inspector/
│   ├── portscan/
│   ├── netwatch/
│   ├── route-inspector/
│   └── ovsdb-inspector/
├── internal/
│   ├── netutil/
│   ├── scanner/
│   ├── monitor/
│   └── config/
└── .github/
    └── workflows/
        └── go.yml
```

### مسئولیت پوشه‌ها

- `docs/`: مستندات مفهومی، معماری، دیاگرام و یادداشت‌های فنی.
- `lessons/`: درس‌های کوچک و تمرین‌ها به ترتیب آموزش.
- `cmd/`: نقطه ورود برنامه‌های اجرایی مستقل.
- `internal/`: کد مشترک داخلی مخزن که قرار نیست به‌عنوان API عمومی کتابخانه ارائه شود.
- `tests/`: در Go الزاماً پوشه مرکزی لازم نیست؛ تست‌ها معمولاً در کنار Package و در فایل‌های `*_test.go` قرار می‌گیرند.
- `.github/workflows/`: گردش کارهای CI مانند Format، Test و Build.

**نکته Module:** در حالت عادی یک `go.mod` در ریشه مخزن کافی است. تا وقتی دلیل فنی مشخص نداریم برای هر درس یا پروژه یک Module جدا نسازیم. مسیر Module پس از مشخص شدن نام کاربری GitHub و URL نهایی مخزن تنظیم می‌شود.

### پیشنهاد اولیه برای README مخزن

README ریشه باید شامل هدف پروژه، وضعیت فعلی دوره، پیش‌نیازها، نسخه Go آزمایش‌شده، نحوه اجرای تست‌ها، فهرست درس‌ها، پروژه‌ها و لینک همین Roadmap باشد.

---

## 7. منابع رسمی و مطالعاتی

### مستندات و آموزش رسمی Go

- [Go Documentation](https://go.dev/doc/) — مستندات رسمی زبان، ابزارها و شیوه استفاده.
- [A Tour of Go](https://go.dev/tour/) — آموزش تعاملی Syntax، ساختار داده و Concurrency.
- [Effective Go](https://go.dev/doc/effective_go) — راهنمای سبک نگارش Go.
- [Go Specification](https://go.dev/ref/spec) — مرجع دقیق قواعد زبان.
- [Go Standard Library](https://pkg.go.dev/std) — مستندات Packageهای استاندارد.
- [net Package](https://pkg.go.dev/net) — TCP، UDP، DNS و Socket.
- [net/http](https://pkg.go.dev/net/http) — HTTP Client و Server.
- [context](https://pkg.go.dev/context) — Cancellation، Deadline و Timeout.
- [testing](https://pkg.go.dev/testing) — تست، Benchmark و Fuzzing.
- [Go Release History](https://go.dev/doc/devel/release) — بررسی تغییرات و اصلاحیه‌های Go.

### کتاب اصلی شبکه

**Network Programming with Go — Adam Woodbeck, No Starch Press (2021)**

- [صفحه ناشر](https://nostarch.com/networkprogrammingwithgo)
- [فهرست مطالب در O'Reilly](https://www.oreilly.com/library/view/network-programming-with/9781098128890/toc.xhtml)

این کتاب از منابع محوری بخش Network Programming است و مباحث Network Architecture، Routing، TCP، UDP، Unix Domain Socket، HTTP، TLS، Serialization، Logging و Metrics را پوشش می‌دهد. چون کتاب مربوط به ۲۰۲۱ است، برای APIهای جدید Go و تغییرات نسخه نصب‌شده، مستندات رسمی جاری Go مرجع نهایی خواهد بود.

### کتاب و دوره تست‌محور

- [Learn Go with Tests — Chris James](https://github.com/quii/learn-go-with-tests) — یادگیری مرحله‌ای Go با رویکرد Test-Driven Development.

### مخازن مهم GitHub برای مطالعه کد واقعی

- [vishvananda/netlink](https://github.com/vishvananda/netlink) — کتابخانه Netlink برای تعامل User Space با Kernel لینوکس.
- [ovn-kubernetes/libovsdb](https://github.com/ovn-kubernetes/libovsdb) — کتابخانه Go برای OVSDB.
- [osrg/gobgp](https://github.com/osrg/gobgp) — پیاده‌سازی BGP با Go؛ برای مطالعه پیشرفته Routing و Control Plane.
- [google/gopacket](https://github.com/google/gopacket) — Packet Capture و Decode و ابزارهای پردازش بسته.

این مخازن از ابتدای دوره برای کپی‌کردن کد استفاده نمی‌شوند. ابتدا مفاهیم و پیش‌نیازها را یاد می‌گیریم، سپس ساختار، تست‌ها، APIها و شیوه طراحی آن‌ها را مطالعه می‌کنیم.

### RFCها و استانداردهای شبکه

- [RFC 9293 — Transmission Control Protocol (TCP)](https://www.rfc-editor.org/rfc/rfc9293)
- [RFC 768 — User Datagram Protocol (UDP)](https://www.rfc-editor.org/rfc/rfc768)
- [RFC 1034 — Domain Names: Concepts and Facilities](https://www.rfc-editor.org/rfc/rfc1034)
- [RFC 1035 — Domain Names: Implementation and Specification](https://www.rfc-editor.org/rfc/rfc1035)
- [RFC 9846 — The Transport Layer Security (TLS) Protocol Version 1.3](https://www.rfc-editor.org/rfc/rfc9846) — نسخه منتشرشده در ژوئیه ۲۰۲۶ که RFC 8446 را جایگزین کرده است.
- [RFC 7047 — The Open vSwitch Database Management Protocol](https://www.rfc-editor.org/rfc/rfc7047)

RFCها را به‌صورت هدفمند و هم‌زمان با درس مرتبط می‌خوانیم؛ لازم نیست همه اسناد از ابتدا تا انتها حفظ شوند.

---

## 8. معیارهای کیفیت کد و پایان هر مرحله

هر پروژه باید تا حد متناسب با سطح آن این موارد را رعایت کند:

- نام‌گذاری روشن و قالب‌بندی با `gofmt`
- مدیریت صریح Error و پرهیز از نادیده گرفتن خطا
- تفکیک منطق، ورودی/خروجی و تنظیمات
- بستن منابع و مدیریت Timeout و Cancellation
- تست حالت عادی و حالت خطا
- عدم قراردادن Secret، Private Key یا داده حساس در مخزن
- README کوتاه شامل هدف، پیش‌نیاز، اجرا و نمونه خروجی
- Commitهای کوچک با پیام معنادار
- استفاده از ابزارهای `go test` و `go vet` و در مباحث Concurrency از Race Detector

### الگوی پیشنهادی Commit

```text
lesson-01: add network information cli
lesson-02: model network interfaces with structs
lesson-03: add validation and error handling
project-tcp: implement echo server
```

---

## 9. وضعیت پیشرفت دوره

این جدول در طول آموزش به‌روزرسانی می‌شود:

| مرحله | وضعیت | تاریخ تکمیل | یادداشت |
|---|---|---|---|
| نصب Go 1.27.4 روی Ubuntu | انجام شده | 2026-11-20 | نصب در `/usr/local/go` و تنظیم PATH در Zsh |
| ساخت اولین برنامه Go | در حال انجام | — | متغیرهای Host، Port و Protocol و خروجی متنی |
| هفته ۱ | شروع شده | — | مبانی اولیه Syntax و ساختار برنامه |
| هفته‌های ۲ تا ۲۱ | شروع نشده | — | طبق Roadmap پیش می‌رویم |

---

## 10. روش استفاده از این سند

1. این فایل را در ریشه مخزن GitHub با نام `ROADMAP.md` قرار بده.
2. لینک آن را از README اصلی مخزن اضافه کن.
3. در پایان هر هفته، وضعیت پیشرفت و پروژه مربوط به آن را ثبت کن.
4. ترتیب پیش‌نیازها را حفظ کن؛ اگر مفهومی هنوز جا نیفتاده، پیش از ورود به مبحث بعدی تمرین تکمیلی انجام بده.
5. تغییرات نسخه Go، کتابخانه‌ها و استانداردها را با مراجعه به مستندات رسمی بررسی کن.

**اصل راهنما:** هدف، تولید کد زیاد نیست؛ هدف، فهمیدن مسئله، انتخاب ابزار مناسب، پیاده‌سازی درست، آزمایش رفتار و توانایی توضیح دادن تصمیم‌های فنی است.
