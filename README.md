# Practice-Go

### Practice golang
- Basic (Examle , Practice)
- Data Structure
- Algorithm
- Goroutine
- Channel
- MongoDB

---
### Goroutine
<aside>
💡 Go 언어를 쓰는 강력한 이유 중에 하나인 고루틴에 대해서 알아보기

</aside>

## 사전지식

### Context Switching

<aside>
💡 불필요한 Context Switching 을 줄이는 것이 성능적으로 매우 중요

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

## 고루틴(Goroutine)

- Go 언어가 사용하는 경량 쓰레드
- Golang에서 실행되는 모든 프로그램은 고루틴에서 실행된다. (메인 함수 등등)
    - Golang의 모든 프로그램은 반드시 하나 이상의 고루틴을 가지고 있음을 의미

## 고루틴 동작 원리

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
  time.Sleep(3 * time.Second) // 왜 있을까??
}
```

- 실행시켜보면 순서대로 PrintAlphabet → PrintNumbers 순으로 실행된다.
- 단, go 키워드를 사용하여 호출한 함수들은 순서와 상관없이 실행된다.
- time.Sleep(3 * time.Second) 은 main 함수에서 3초간 대기하는 코드이다.
    - 왜 필요할까?
    - 이 time 코드가 없다면, main 함수는 go키워드로 실행한 함수의 결과를 기다리지 않는다. 따라서 1,2 번의 함수만 실행되고 메인 고루틴이 종료되버리는데, 그러면 go키워드로 실행한 서브 고루틴도 함게 종료되므로 어떠한 결과도 출력되지 않게 된다.
