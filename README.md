# GO

GO 공식 사이트: https://go.dev/ (혹은 https://golang.org)

- GO 프로그래밍 언어는 2007년 구글에서 개발을 시작하여 2012년 GO 버젼 1.0을 완성하였다. GO는 이후 계속 향상된 버젼을 내 놓았으며 2022년 초에는 버젼 1.18 에 이르렀다.
흔히 golang 이라고도 불리우는 Go 프로그래밍 언어는 구글의 V8 Javascript 엔진 개발에 참여했던 Robert Griesemer, Bell Labs에서 유닉스 개발에 참여했던 Rob Pike, 그리고 역시 Bell Labs에서 유닉스 개발했으며 C 언어의 전신인 B 언어를 개발했던 Ken Thompson이 함께 개발하였다.

-  key features :
    - 컴파일 언어 compile to standalone binaries (application을 실행하기 위한 모든 파일, 라이브러리 등이 바이너리로 저장됨)
    - 정적 타입 (statically typed), strong and statically typed (한번 type이 선언되면 변하지 않음)
    - garbage collection 기능 제공
    - CSP(communicating sequential processes) 스타일의 concurrent 프로그래밍을 지원
    - simplicity
    - fast compile times
    - garbage collected => not managing my own memory
    - built-in concurrency

---

https://www.youtube.com/watch?v=YS4e4q9oBaU

⭐️ Course Contents ⭐️  

✓ (0:00:00) Introduction  

✓ (0:16:57) Setting Up a Development Environment  

✓ (0:35:48) Variables 변수
  - variable declaration  
  - redeclaration and shadowing  
  - all variables must be used
  - visibility 
    - lower case first letter for package scope
    - upper case first letter to export
    - no private scope
  - naming conventions  
  - type conversions    

✓ (0:57:05) Primitives 데이터 타입, 아주 기초의 것  
  - boolean type
    - values are true or false
    - not an alias for other types(e.g. int)
    - zero value is false
  - numeric types (integers, floating point, complex numbers)
    - integers 
      - signed integers (value can be +/-)
        - int type has varing size, bun min 32 bits
        - 8 bit(int8) through 64 bit (int64)
      - unsigned integers (value +)
        - 8 bit(byte and unit8) through 32 bit (unit32)
      - arithmetic operations
        - addition, substraction, multiplication, division, remainder
      - bitwise operations
        - and, or, xor, not
      - zero value is 0 
      - can't mix types in same family (unit16 + unit32 = error)
    - floating point numbers
      - follow IEEE-754 standard
      - zero value is 0
      - 32 and 64 bit versions
      - Literal styles 
        - decimal (3.14)
        - exponential (13e18 or 2E10)
        - mixed (13.7e12)
      - arithmetic operations
        - addition, substraction, multiplication, division
    - complex numbers
      - zero value is (0+0i)
      - 64 and 128 bit versions (real + imagine )
      - built-in functions
        - complex : make complex number from two floats
        - real : get real part as float
        - imag : get imaginary part as float
      - arithmetic operations
          - addition, substraction, multiplication, division
  - text types (strings, rune)
    - strings
      - UTF-8
      - immutable (불변)
      - can be concatenated with + operator
      - can be converted to []byte <-> string
    - rune
      - UTF-32
      - alias for int32
      - special methods normally required to process
        - e.g. strings.Reader#ReadRune

️✓ (1:26:29) Constants 상수 <-> 변수 variable   
  - constants 
    - immutable, but can be shadowed
    - replaced by the compiler at compile time
      - value must be calculable at compile time
  - naming convention
    - named like variables
      - PascalCase for exported constants
      - camelCase for internal constants
  - typed constants
    - typed constants work like immutable variables
      - can interoperate only with same type
  - untyped constants
    - untyped constants work like literals
      - can interoperate with similar types
  - enumerated constants
    - special symbol iota allows related constants to be created easily
    - iota starts at 0 in each const block and increments by one
    - watch out of constant values that match zero values for variables
  - enumeration expressions
    - operations that can be determined at compile time are allowed
      - arithmetic
      - bitwise operations
      - bitshifting (<<, >>)

⌨️ (1:47:53) Arrays and Slices  

⌨️ (2:17:20) Maps and Structs  

⌨️ (2:48:00) If and Switch Statements  

⌨️ (3:21:17) Looping  

⌨️ (3:41:34) Defer, Panic, and Recover  

⌨️ (4:03:57) Pointers

⌨️ (4:21:30) Functions  

⌨️ (4:57:59) Interfaces  

⌨️ (5:33:57) Goroutines  

⌨️ (6:05:10) Channels  

---

- strong and statically typed (한번 type이 선언되면 변하지 않음)
- excellent community
- key features :
    - simplicity
    - fast compile times
    - garbage collected => not managing my own memory
    - built-in concurrency
    - compile to standalone binaries (application을 실행하기 위한 모든 파일, 라이브러리 등이 바이너리로 저장됨)