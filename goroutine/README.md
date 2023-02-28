## Goroutine
<aside>
💡 Go 언어를 쓰는 강력한 이유 중에 하나인 고루틴에 대해서 알아보기

</aside>

### 사전지식

### Context Switching

<aside>
- 불필요한 Context Switching 을 줄이는 것이 성능적으로 매우 중요

</aside>

- **Context** : 프로세스와 관련된 정보들의 집합
    - CPU register context : CPU 안에 저장되어 있는 Context
    - Code & Data, Stack PCB : Memory 안에 저장되어 있는 Context  
  

- Context Saving : 현재 프로세스의 Register Context를 저장하는 작업  


- Context Restoring : Register Context를 프로세스로 복구하는 작업  


- **Context Switching  : 현재 진행하고 있는 Task(Process, Thread)의 상태를 저장하고 다음 진행할 Task의 상태 값을 읽어 적용하는 과정**
    - 실행 중인 프로세스의 Context는 저장
    - 앞으로 실행 할 프로세스의 Context는 복구
    - 커널의 개입으로 이루어짐

### 쓰레드(Thread)

- 프로세서 활용의 기본 단위
- 제어 요소 외 코드, 데이터 및 자원들은 프로세스의 다른 쓰레드들과 공유
- 하나의 CPU가 여러 쓰레드를 다룰 때 쓰레드를 전환시키며 CPU를 사용하도록 한다. (Context Switching)
    - CPU 개수와 쓰레드의 개수가 동일하면 Context Switching이 발생하지 않는다.

---

### 고루틴(Goroutine)

- Go 언어가 사용하는 경량 쓰레드
- Golang에서 실행되는 모든 프로그램은 고루틴에서 실행된다. (메인 함수 등등)
    - Golang의 모든 프로그램은 반드시 하나 이상의 고루틴을 가지고 있음을 의미

### 고루틴 동작 원리

- 고루틴은 OS 쓰레드를 이용하는 경량 쓰레드(lightweight thread) 이다.
    - 자바 쓰레드는 커널 레벨 쓰레드
- 고루틴을 아무리 많이 생성해도 OS 쓰레드를 이용하기 때문에 OS 레벨의 Context Switching 비용이 발생하진 않음

### 단순 고루틴 확인 예제(ex1)

```go
// 함수는 ex1에서 참고

func main() {
  PrintAlphabet() // 1
  fmt.Println("")

  PrintNumbers() // 2
  fmt.Println("")

  go PrintAlphabet()
  go PrintNumbers()
  time.Sleep(3 * time.Second) // 왜 있을까?
}
```

- 실행시켜보면 순서대로 PrintAlphabet → PrintNumbers 순으로 실행된다.
- 단, go 키워드를 사용하여 호출한 함수들은 순서와 상관없이 실행된다.
- time.Sleep(3 * time.Second) 은 main 함수에서 3초간 대기하는 코드이다.
    - 왜 필요할까?
    - 이 time 코드가 없다면, main 함수는 go키워드로 실행한 함수의 결과를 기다리지 않는다. 따라서 1,2 번의 함수만 실행되고 메인 고루틴이 종료되버리는데, 그러면 go키워드로 실행한 서브 고루틴도 함게 종료되므로 어떠한 결과도 출력되지 않게 된다.


---

### EX2 - 서브 고루틴 대기

Ex1 처럼 메인 고루틴이 종료되면 서브 고루틴도 한꺼번에 종료되는 데 이를 어떻게 해결할까?

```go
var wg sync.WaitGroup

wg.Add(2)

wg.Done()
wg.Wait()
```

- Golang 에서는 서브 고루틴이 종료될 때까지 메인 고루틴이 유지될 수 있도록 WaitGroup을 제공
- WaitGroup의 Add를 통해 메인 고루틴이 대기할 서브 고루틴의 수를 정의
- 서브 고루틴에서는 Done()을 통해 고루틴의 종료를 알린다.
- 메인 고루틴에서는 Wait()를 통해 모든 서브 고루틴이 종료될 때 까지 대기

### go func() & go func(i int)

> **go func()** : 매개변수가 없는 함수를 고루틴으로 실행할 때 사용하는 구문 (like lambda)
>
> 고루틴 내에서 전역 변수나 외부 변수를 참조할 수 있으므로, **공유 변수에 접근하는 경우** 경쟁 상태와 같은 문제가 발생할 수 있음

> **go func(i int)** : 매개변수 i를 받는 함수를 고루틴으로 실행할 때 사용하는 구문
> 
> 이 경우 각 고루틴은 자신이 실행하는 동안 매개변수 i의 값을 참조하게 되며, 고루틴은 각각 서로 다른 값을 참조할 수 있음
> 
> 이를 통해 각 고루틴이 개별적으로 작업을 수행할 수 있고, 고루틴이 실행하는 동안 전역 변수나 외부 변수에 대한 접근이 제한되므로 경쟁 상태와 같은 문제가 발생하지 않음

#### Example
```go
func main() {
	for i := 1; i <= 3; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(1 * time.Second)
}
// result
4
4
4
```
- 고루틴은 for loop 에 의해 **총 3개**가 만들어진다. 주의할 점은 고루틴의 로직이 다 실행되고 i 값이 증가하는 것은 아니다. (기존 for loop 와의 차이)
- for loop 에서의 i 값이 고루틴에서 출력되는 시점에는 이미 4가 되어 있기 때문이다.
- 모든 고루틴이 동시에 i값을 참조하고 있기 때문의, 각각의 고루틴(여기선 3개)이 동시에 i 값이 4인 상태에서 실행되어 4가 출력된다.

```go
func main() {
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(1 * time.Second)
}

// result
1
2
3
```
- 위와 다르게 `i` 라는 변수를 각각의 고루틴이 개별적으로 가지고 있다. 
- 각 고루틴에서는 개별적으로 전달받은 `i` 값을 출력하게 되므로 `1,2,3` 의 값을 반환한다. 단, 고루틴은 병렬 처리를 하므로 실행 순서가 보장되지 않아 `2,1,3` 같은 순서도 반환될 가능성이 있다.
> 근데 왜 `go func(i)` 가 아니라 다른 변수처럼 go func(i int) 라고 적는 것일까?
> 
> 다른 예시로 살펴보면,

```go
for i := 0; i < workers; i++ {
		go func(workerNum int) {
			
			// ... 생략
			
			}
		}(i)
	}
```
- 주의할 점은 workerNum 은 i 의 값을 익명함수로 전달해주는 역할을 한다.
- 즉 위에서도 i 와 익명 함수 안에 있는 i는 같은 것이 아니며, 단순히 for loop 의 i 값을 익명함수로 전달해주는 것이다.
- 이 예제에서 만약 workers = 3 이라면, workerNum 에도 `0,1,2` 가 차례대로 전달된다.

### Lock
- 동시성(Concurrency)을 다룰 때 발생할 수 있는 문제 중 하나로, 공유 자원에 대한 접근을 제어하는 동기화 매커니즘
- 여러 개의 쓰레드나 고루틴이 동시에 하나의 자원에 접근하려고 할 때, 일어날 수 있는 문제들 중 하나가 데이터 무결성(Data Integrity)이 깨지는 것
  - 예를 들어, 한 쓰레드가 공유 자원의 값을 변경하고 있을 때, 다른 쓰레드가 그 값을 읽을 경우, 변경이 완료되기 전의 값이 읽혀지는 등의 문제 발생 가능
> 이러한 문제들을 해결하기 위해 Lock 을 사용한다.
- Lock 은 공유 자원에 대한 접근을 동기화하기 위한 가장 대표적인 매커니즘으로, Lock 을 사용하면 한 번에 하나의 쓰레드나 고루틴만이 공유 자원에 접근할 수 있도록 제어할 수 있다.