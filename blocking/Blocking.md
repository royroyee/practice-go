# Blocking , Non-Blocking

### 사전지식
- `제어권` : 자신(함수)의 코드를 실행할 권리 같은 것. 제어권을 가진 함수는 자신의 코드를 끝까지 실행한 후, 자신을 호출한 함수에게 돌려준다.
- `결과값을 기다리는 것` : A 함수에서 B 함수를 호출했을 때, A 함수가 B 함수의 결과값을 기다리는 여부

### 동기 Synchronous & 비동기 Asynchronous
- 처리해야 할 작업들을 어떠한 흐름으로 처리할 것인가
- 호출되는 함수의 작업 완료 여부를 신경쓰는 지, 함수 실행 및 리턴 의 순차적인 흐름을 따르는 지 
> Synchronous : 작업을 동시에 수행하거나, 동시에 끝나거나, 끝나는 동시에 시작
> 
> Asynchronous : 시작, 종료가 일치하지 않으며, 끝나는 동시에 시작을 하지 않음

### 동기 Synchronous
- 호출하는 함수 A 가 호출되는 함수 B의 작업 완료 후 리턴을 기다리거나, 바로 리턴 받더라도 미완료 상태라면 작업 완료 여부를 스스로 계속 확인하며 신경쓰는 것
  - 함수 A가 함수B를 호출한 뒤에도 함수 B의 리턴 값을 계속 확인하면서 신경쓰는 것

### 비동기 Asynchronous
- 함수 A가 함수 B를 호출할 때 콜백 함수를 함께 전달해서, 함수 B의 작업이 완료되면 함께 보낸 콜백 함수를 실행
  - 함수 A는 함수 B를 호출한 후로 함수 B의 작업 완료 여부에는 신경 X 

### Blocking & Non-Blocking
- 처리되어야 하는 작업이 전체적인 작업 흐름을 막는지 안 막는지에 대한 것
  - 즉 **제어권**이 누구한테 있는 지
> Blocking : 자신의 작업을 진행하다가 다른 주체의 작업이 시작되면 다른 작업이 끝날 때까지 기다렸다가 자신의 작업을 시작
> Non-Blocking : 다른 주체의 작업에 관련없이 자신의 작업을 하는 것

### Blocking 
- A 함수가 B 함수를 호출하면, 제어권을 A가 호출한 B함수에 넘겨준다.
  1. A 함수가 B 함수를 호출하면 B에게 제어권을 넘긴다.
  2. 제어권을 넘겨받은 B는 함수를 실행하고, A는 B에게 제어권을 넘겨주었기 때문에 실행을 멈춘다.
  3. B함수가 실행이 끝나면, 자신을 호출한 A에게 제어권을 돌려준다.
  4. A함수는 나머지 로직을 실행한다. (다시 함수 실행)
- 가장 일반적인 방식 

### Non-Blocking
- A 함수가 B 함수를 호출해도 제어권은 그대로 자신이 갖고 있는다.
  1. A 함수가 B 함수를 호출
  2. B 함수는 실행되지만, 제어권은 A 함수가 그대로 가지고 있다.
  3. 즉, A 함수는 B 함수를 호출한 이후에도(B함수가 실행되는 중에도) 자신의 코드를 계속 실행한다.

### Example (blocking.go)
```go
func blockingOperation() {
	fmt.Println("start blocking operation...")
	time.Sleep(time.Second * 3)
	fmt.Println("end blocking operation!")
}

func nonBlockingOperation(c chan string) {
	fmt.Println("start non-blocking operation...")
	time.Sleep(time.Second * 3)
	c <- "end non-blocking operation!"
}

func main() {
	// 블로킹 작업 수행
	blockingOperation()

    // non-blocking 작업 결과를 기다리지 않고 다른 작업 수행 가능 (여기선 nonBlockingOperation 함수에 따라 약 3초 동안 수행)
    fmt.Println("start non-blocking operation...(main)")
	c := make(chan string)
	go nonBlockingOperation(c)

	// 논블로킹 작업 결과를 기다리지 않고 다른 작업 수행
	fmt.Println("do other work while waiting for non-blocking operation...")

	// 논블로킹 작업 결과 출력
	fmt.Println(<-c)
}

// result
start blocking operation...
end blocking operation
start non-blocking operation...(main)
start non-blocking operation...
do other work while waiting for non-blocking operation
end non-blocking operatinon!

```
- `blockingOperation()` : 3초 동안 대기하는 blocking 작업을 수행한다. 이 함수가 수행되는 동안 다른 작업은 수행 불가능
- `nonBlockingOperation()` : 3초 동안 대기하는 non-blocking 작업을 수행한다. channel 을 이용하여 다른 goroutine 에서 수행되기 때문에 이 함수가 수행되는 동안 다른 작업 수행 가능
  - `do other ~` 문장이 `end non-blocking~` 보다 먼저 출력된 것으로 확인 가능

## 조합

### Sync & Blocking
- A 함수는 B 함수의 리턴 값을 필요로 한다. (Synchronous)
- 제어권을 B에게 넘겨주고 B 함수의 실행이 완료되어 리턴 값과 제어권을 돌려줄 때 까지 기다린다. (Blocking)
> 요청받은 함수의 작업이 끝나야 제어권을 돌려받음 + 요청 함수는 결과가 나올때까지 계속 지켜봄

### Sync & Non-Blocking
- A 함수는 B 함수를 호출하는데, A 함수는 B 함수에게 제어권을 주지 않고 자신의 코드를 계속 실행한다. (Non-blocking)
- 하지만 A 함수는 B 함수의 리턴 값이 필요하기 때문에, 중간중간 B 함수에게 함수 실행을 완료했는 지 물어본다. (Synchronous)
> 제어권은 바로 돌려주지만 요청 함수는 결과가 나올 때 까지 계속 확인

### Async-NonBlocking
- A 함수는 B 함수를 호출한다. 이 때, 제어권을 B 함수에게 주지 않고 자신이 계속 가지고 있는다.
- 따라서 B 함수를 호출한 이후에도 멈추지 않고 자신의 함수를 계속 실행한다. (Non-Blocking)
- 그리고 B 함수를 호출할 때 콜백함수를 함께 준다. B함수는 자신의 작업이 끝나면 A 함수가 준 콜백함수를 실행한다. (비동기)

### Async-Blocking
- (잘 볼 수 없는 조합)
- A 함수는 B 함수의 리턴 값에 신경쓰지 않고 콜백함수를 보낸다. (비동기)
- B 함수의 작업에 관심 없음에도 불구하고 A 함수는 B 함수에게 제어권을 넘긴다. (Blocking)
- 따라서 A 함수는 자신과 관련 없는 B 함수의 작업이 끝날 때 까지 기다려야 한다.

