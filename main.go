package main

import (
	"fmt"
	"sync"
)

//func main() {
//	var fastOrderSlice []int
//	fastOrder := []int{6,7,13,8,9,11,10, 3, 1,2,5, 4}
//sortingSlice:=getSortingSlice(fastOrder)
//fmt.Println(sortingSlice)
//fastOrderSlice=append(fastOrderSlice, sortingSlice[0])
//}
//func getSortingSlice(fastOrder []int)[]int  {
//
//
//	var first []int
//	var second []int
//
//	if len(fastOrder)<2{
//		return fastOrder
//	}
//	for i:=1; i<len(fastOrder); i++ {
//		if fastOrder[i] > fastOrder[0] {
//			second = append(
//				second,
//				fastOrder[i]			)
//
//		} else {
//			first = append(
//				first,
//				fastOrder[i]			)
//
//		}
//	}
//less:=getSortingSlice(first)
//last:=getSortingSlice(second)
//less=append(less, fastOrder[0])
//less=append(less, last[0:]...)
//	return less
////}
//func main()  {
//	orderSlice:=[]int{5,6,1,2,4,3,1,9,8,4,6}
//	result:=GetOrderSlice(orderSlice)
//fmt.Println("Результат", result)
//}
//
//func GetOrderSlice(orderSlice[]int)([]int)  {
//	//var result [] int
//	var left []int
//	var right [] int
//	if len(orderSlice)<2{
//		return orderSlice
//	}
//	lenHalf:=len(orderSlice)/2
//	left= orderSlice[:lenHalf]
//	right=orderSlice[lenHalf:]
//	left=GetOrderSlice(left)
//	right=GetOrderSlice(right)
//	fmt.Println("left",left)
//	fmt.Println("right", right)
//	fmt.Println(orderSlice)
//return getSort(left, right)
//}
//
//func getSort(left[] int, right[]int)[]int  {
//	var result [] int
//	i, j := 0, 0
//	for i < len(left) && j < len(right) {
//		if left[i] < right[j] {
//			result = append(result, left[i])
//			i++
//		} else {
//			result = append(result, right[j])
//			j++
//		}
//	}
//
//	result = append(result, left[i:]...)
//	result = append(result, right[j:]...)
//
//	return result
//}

//type ListNode struct {
//	val int
//	node *ListNode
//}
//
//func main()  {
//	l1:=[]int{2,3,4,5,6,7}
//	l2:=[]int{1,2,4,5,7,8,9}
//	list1:=GetStruct(l1,l2)
//	fmt.Print(list1)
//
//	print(list1)
//}
//
//func print(l*ListNode)  {
//	current:=l
//	for current!=nil{
//		fmt.Println(current.val)
//		current=current.node
//	}
//}
//func GetStruct(l []int, l2 []int )*ListNode{
//	head:=&ListNode{0, nil}
//	current:=head
//	for _, val2:=range l2 {
//		list1 := &ListNode{val: val2, node: nil}
//		current.node = list1
//		current = current.node
//	}
//		for _, val := range l {
//			list1 := &ListNode{val: val, node: nil}
//			current.node = list1
//			current = current.node
//		}
//
//
//	return head.node
//}

//type Square struct{
//	а int
//
//}
//type Circle struct {
//	radius int
//}
//
//type SquareFiggures interface {
//	GetSquareFiggures (a int) int
//
//}
//func main()  {
//
//var v, ok:= SquareFiggures
//fmt.Println("Площадь квадрата", v.GetSquareFiggures(2))
//
//var c SquareFiggures=&Circle{}
//	fmt.Println("Площадь круга", c.GetSquareFiggures(10))
//}
//func (s*Square)GetSquareFiggures(a interface{})   {
//	fmt.Println("a",a)
//
//}
//
//func (c*Circle) GetSquareFiggures (radius int)int{
//	square:=math.Pi*float64(radius)*float64(radius)
//	return int(square)
//}
//Задача на написание генератора случайных чисел
//func writeToChannel(ch chan<- int) {
//	for i := 0; i < 5; i++ {
//		ch <- i
//	}
//	close(ch)
//}
//
//func readFromChannel(ch <-chan int) {
//	for value := range ch {
//		fmt.Println(value)
//	}
//}
//
//func main() {
//	ch := make(chan int)
//	go writeToChannel(ch)
//	readFromChannel(ch)
//}

//func randNumsGenerator(n int) <-chan int {
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//
//	out := make(chan int)
//	uniqueNums := make(map[int]bool)
//	go func() {
//		for i := 0; i < n; i++ {
//			num := r.Intn(20)
//			for uniqueNums[num] {
//				num = r.Intn(20)
//			}
//			uniqueNums[num] = true
//			out <- num
//		}
//		close(out)
//	}()
//	return out
//}
//
//
//func main() {
//	var result []int
//	for num := range randNumsGenerator(10) {
//		fmt.Println(num)
//		result=append(result, num)
//	}
//	fmt.Println(result)
//}
//Даны два канала. В первый пишутся числа. Нужно, чтобы числа читались из первого по мере поступления,
//что-то с ними происходило (допустим, возводились в квадрат) и результат записывался во второй канал.

//func main() {
//	chanel1 := make(chan int)
//	chanel2 := make(chan int)
//	go writeToChanel(chanel1)
//	go readchanal(chanel1, chanel2)
//	for n := 0; n < 10; n++ {
//		num := <-chanel2
//		fmt.Println("Данные записанные в канал", num)
//	}
//}
//func writeToChanel(chanel chan int) {
//	for {
//		num := rand.Intn(100)
//		chanel <- num
//		time.Sleep(time.Millisecond * 500)
//	}
//	close(chanel)
//}
//func readchanal(chanel1 chan int, chanel2 chan int) <-chan int {
//	for v := range chanel1 {
//		chanel2 <- v
//	}
//	close(chanel1)
//	close(chanel2)
//	return chanel2
//}

//func main(){
//	num:=3
//	chanel1:=make(chan int)
//	chanel2:=make(chan int)
//	chanel3:=make(chan int)
//
//
//		go writeChanel(chanel1)
//		//
//		go writeChanel(chanel2)
//		go writeChanel(chanel3)
//	for n:=0; n<num; n++ {
////	num:=<-chanel1
////	fmt.Println("Первый канал",num)
////	num=<-chanel2
////		fmt.Println("Второй канал",num)
////	num=<-chanel3
////		fmt.Println("Третий канал",num)
////fmt.Println(num)
//
//	for num:=range joinChanel(chanel1,chanel2,chanel3){
//fmt.Println(num)
//	}
//}
//}
//
//func writeChanel(chanel chan int) {
//	for  {
//		num:=rand.Intn(100)
//		chanel<-num
//		time.Sleep(time.Millisecond*200)
//	}
//	close (chanel)
//}
//
//func joinChanel(chs ...<-chan int) <-chan int{
//	mergedCh := make(chan int)
//
//	go func() {
//		wg := &sync.WaitGroup{}
//
//		wg.Add(len(chs))
//
//		for _, ch := range chs {
//			go func(ch <-chan int, wg *sync.WaitGroup) {
//				defer wg.Done()
//				for id := range ch {
//					mergedCh <- id
//				}
//			}(ch, wg)
//		}
//
//		wg.Wait()
//		close(mergedCh)
//	}()
//
//	return mergedCh
//}
//func main() {
//	var wg sync.WaitGroup
//	c := make(chan string)
//	s := "Hello word"
//	s = s + " "
//
//	for n := 0; n < 5; n++ {
//		wg.Add(1)
//		go myFunc(&wg, s, c)
//	}
//	for v := range c {
//		fmt.Println(v)
//	}
//	wg.Wait()
//}
//
//func myFunc(wg *sync.WaitGroup, s string, c chan string) <-chan string {
//	c <- s
//	time.Sleep(time.Millisecond * 500)
//	close(c)
//	wg.Done()
//	return c
//}
//func main() {
//
//	c := make(chan string)
//	s := "Hello word"
//	s = s + " !"
//
//	for n := 0; n < 5; n++ {
//
//		go myFunc(s, c)
//	}
//	for v := range c {
//		fmt.Println(v)
//	}
//
//}
//
//func myFunc(s string, c chan string) <-chan string {
//	c <- s
//	time.Sleep(time.Millisecond * 500)
//	close(c)
//
//	return c
//}
//func main() {
//
//	inputChanel := make(chan string)
//	outputChanel := make(chan string)
//	nums := "112234567"
//
//	removeDuplicates(inputChanel, outputChanel, nums)
//	for v := range outputChanel {
//		fmt.Println(v)
//	}
//}
//
//func removeDuplicates(inputChanel chan string, outputChanel chan string, num string) {
//	go func() {
//		inputChanel <- num
//		close(inputChanel)
//	}()
//	go func() {
//		nums := <-inputChanel
//
//		var s string
//		prev := nums[0]
//		s = s + string(prev)
//		for i := 1; i < len(nums); i++ {
//			if nums[i] != prev {
//				s = s + string(nums[i])
//				prev = nums[i]
//			}
//		}
//
//		outputChanel <- s
//		close(outputChanel)
//
//	}()
//}

//
//func main() {
//	firstChan := make(chan int)
//	secontChan := make(chan int)
//	stopChan := make(chan struct{})
//	var num = []int{1, 2, 3, 4}
//
//	for i := 0; i < len(num); i++ {
//		go func(i int) {
//			firstChan <- num[i]
//		}(i)
//		go func(i int) {
//			secontChan <- num[i]
//		}(i)
//	}
//	go func() {
//		// Отправляем сигнал остановки через 1 секунду
//		time.Sleep(time.Second)
//		close(stopChan)
//	}()
//	for v := range calculator(firstChan, secontChan, stopChan) {
//		fmt.Println(v)
//	}
//}
//func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
//
//	thirdChan := make(chan int)
//
//	go func() {
//		defer close(thirdChan)
//		for {
//			select {
//			case <-firstChan:
//				nums := <-firstChan
//				thirdChan <- nums * nums
//			case <-secondChan:
//				num := <-secondChan
//				thirdChan <- num
//			case <-stopChan:
//				return
//			}
//		}
//	}()
//
//	return thirdChan
//}
//func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
//	for i := 0; i < n; i++ {
//		x1 := <-in1
//		x2 := <-in2
//		result := fn(x1) + fn(x2)
//		out <- result
//	}
//}
//
//func main() {
//	// пример использования функции merge2Channels
//	sum := func(x int) int { return x + x }
//	in1 := make(chan int)
//	in2 := make(chan int)
//	out := make(chan int)
//
//	go func() {
//		defer close(in1)
//		for i := 1; i <= 5; i++ {
//			in1 <- i
//		}
//	}()
//
//	go func() {
//		defer close(in2)
//		for i := 6; i <= 10; i++ {
//			in2 <- i
//		}
//	}()
//
//	go merge2Channels(sum, in1, in2, out, 5)
//
//	for i := 1; i <= 5; i++ {
//		fmt.Println(<-out)
//	}
//}
//Написать конвейер для возведения чисел в квадрат
//если приходят данные из одного канала число возводится к вадрат если из другого умножается на 3
//иначе done

//func main() {
//	in1 := make(chan int)
//	in2 := make(chan int)
//	out := make(chan int)
//	var wg sync.WaitGroup
//	for n := 0; n < 5; n++ {
//		wg.Add(1)
//		go func(n int) {
//			in1 <- n
//			wg.Done()
//		}(n)
//	}
//	for n := 10; n < 15; n++ {
//		wg.Add(1)
//		go func(n int) {
//			in2 <- n
//			wg.Done()
//		}(n)
//	}
//	go mergeChannels(in1, in2, out)
//	go func() {
//		wg.Wait()
//		close(in1)
//		close(in2)
//	}()
//	for v := range out {
//		fmt.Println(v)
//	}
//}
//
//func mergeChannels(in1, in2, out chan int) {
//	defer close(out)
//	for {
//		select {
//		case n, ok := <-in1:
//			if !ok {
//				for n := range in2 {
//					out <- 3 * n
//				}
//				return
//			}
//			out <- n * n
//		case n, ok := <-in2:
//			if !ok {
//				for n := range in1 {
//					out <- n * n
//				}
//				return
//			}
//			out <- 3 * n
//		}
//	}
//}
//ОЧЕНЬ ХОРОШИЙ ПРИМЕР смержить несколько каналов в 1

//func main() {
//	var a, b, c = make(chan int), make(chan int), make(chan int)
//	go func() {
//		for n := 0; n < 5; n++ {
//			a <- n
//		}
//		close(a)
//	}()
//	go func() {
//		for n := 6; n < 10; n++ {
//			b <- n
//		}
//		close(b)
//	}()
//	go func() {
//		for n := 11; n < 15; n++ {
//			c <- n
//		}
//		close(c)
//	}()
//
//	for num := range joinChannels(a, b, c) {
//		fmt.Println(num)
//	}
//}
//
//func joinChannels(in ...chan int) <-chan int {
//	merge := make(chan int, len(in))
//	done := make(chan struct{})
//	for i := 0; i < len(in); i++ {
//		go func(c chan int) {
//			defer func() { done <- struct{}{} }()
//			for {
//				value, ok := <-c
//				if !ok {
//					return
//				}
//				merge <- value
//			}
//		}(in[i])
//	}
//	go func() {
//		for i := 0; i < len(in); i++ {
//			<-done
//		}
//		close(merge)
//	}()
//	return merge
//}

//конвеер чисел
//func main() {
//	a := make(chan int)
//	b := make(chan int)
//	go func() {
//		for i := 0; i < 5; i++ {
//			a <- i
//		}
//		close(a)
//	}()
//	go conveer(a, b)
//	for v := range b {
//		fmt.Println(v)
//	}
//}
//func conveer(a chan int, b chan int) {
//	for v := range a {
//		b <- v * v
//	}
//	close(b)
//}
//Создать гоурутину которая будет брать по порядку числа от 1 до 10 прибавлять к ним 2 и отдавать в другую гоурутину
//func main() {
//	firstChan := make(chan int)
//	go func() {
//		for i := 0; i < 10; i++ {
//			firstChan <- i + 10
//		}
//		close(firstChan)
//	}()
//	for v := range getSum(firstChan) {
//		fmt.Println(v)
//	}
//}
//
//func getSum(first chan int) chan int {
//	secondChan := make(chan int)
//
//	go func() {
//		for v := range first {
//			secondChan <- v * v
//		}
//		close(secondChan)
//	}()
//
//	return secondChan
//
//}
//func main() {
//	firstChan := make(chan int)
//	done := make(chan struct{})
//	var su sync.RWMutex
//	var wg sync.WaitGroup
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go func(i int) {
//			su.Lock()
//			firstChan <- i
//			su.Unlock()
//			wg.Done()
//		}(i)
//	}
//	go func() {
//		wg.Wait()
//		close(firstChan)
//	}()
//	go func() {
//		time.Sleep(2 * time.Second)
//		close(done)
//	}()
//	for v := range getSum(firstChan, done) {
//		fmt.Println(v)
//	}
//	//for v := range firstChan {
//	//	fmt.Println(v)
//	//}
//}
//
//func getSum(first chan int, done chan struct{}) <-chan int {
//	secondChan := make(chan int)
//
//	go func() {
//		defer close(secondChan)
//		for {
//			select {
//			//очень важная конструкция!!!!
//			case n, ok := <-first:
//				if !ok {
//					return
//				}
//				secondChan <- n * n
//
//			case <-done:
//				secondChan <- 1
//				return
//				//default:
//				//	secondChan <- -1
//
//			}
//		}
//	}()
//	return secondChan
//}
//func countTo(max int) (<-chan int, func()) {
//	ch := make(chan int)
//	done := make(chan struct{})
//	cancel := func() {
//		close(done)
//	}
//	go func() {
//		for i := 0; i < max; i++ {
//			select {
//			case <-done:
//				return
//			default:
//				ch <- i
//			}
//		}
//		close(ch)
//	}()
//	return ch, cancel
//}
//func main() {
//	ch, cancel := countTo(10)
//	for i := range ch {
//		if i > 5 {
//			break
//		}
//		fmt.Println(i)
//	}
//	cancel()
//}
//func main() {
//	firstChan := make(chan int)
//	secondChan := make(chan int)
//	nums := []int{30, 40, 50, 70}
//	nums2 := []int{100, 300, 500, 600}
//	var wg1 sync.WaitGroup
//	for _, num := range nums {
//		wg1.Add(1)
//		go func(n int) {
//			firstChan <- n
//			wg1.Done()
//		}(num)
//	}
//	go func() {
//		wg1.Wait()
//		close(firstChan)
//
//	}()
//	var wg sync.WaitGroup
//	for _, num2 := range nums2 {
//		wg.Add(1)
//		go func(n int) {
//			secondChan <- n
//		}(num2)
//	}
//	go func() {
//		wg.Wait()
//		close(secondChan)
//	}()
//	out, chansel := getNewChanel(firstChan, secondChan)
//	count := 0
//	for v := range out {
//		if count > 6 {
//			chansel()
//			break
//		}
//		fmt.Println(v)
//		count++
//	}
//}
//func getNewChanel(firstChan chan int, secondChan chan int) (chan int, func()) {
//	out := make(chan int)
//	done := make(chan struct{})
//	chansel := func() {
//		close(done)
//	}
//	go func() {
//		defer close(out)
//		for {
//			select {
//			case k, ok := <-firstChan:
//				if !ok {
//					return
//				}
//
//				out <- k / 10
//			case n, ok := <-secondChan:
//				if !ok {
//					return
//				}
//				out <- n * 10
//			case <-done:
//				out <- 0
//				return
//			case <-time.After(1 * time.Millisecond):
//				return
//			}
//		}
//	}()
//	return out, chansel
//}
//func add(w *sync.WaitGroup, num *int32) {
//	defer w.Done()
//	atomic.AddInt32(num, 1)
//}
//func main() {
//	var n int32 = 0
//	var wg sync.WaitGroup
//	wg.Add(1000)
//	for i := 0; i < 1000; i = i + 1 {
//		go add(&wg, &n)
//	} // create 1000 new goroutines
//	wg.Wait()
//	println(n)
//}

//func main() {
//	pool := workerpool.New(5)
//
//	for i := 0; i < 10; i++ {
//		pool.Submit(func() {
//			dur := time.Duration(rand.Intn(3)+1) * time.Second
//			time.Sleep(dur)
//			fmt.Printf("worker completed job %d in %v\n", i, dur)
//		})
//	}
//
//	pool.StopWait()
//	fmt.Println("all jobs completed")
//}
//func main() {
//	//runtime.GOMAXPROCS(1)
//	for i := 0; i < 5; i++ {
//		go func() {
//
//			fmt.Print(i)
//
//		}()
//		runtime.Gosched()
//	}
//	time.Sleep(2 * time.Second)
//}
//func main() {
//
//	num := 38
//	result := IntToRoman(num)
//	fmt.Println(result)
//}
//func IntToRoman(num int) string {
//	//var m = map[int]string{10: "X", 1: "I", 5: "V", 50: "L", 100: "C", 500: "D", 1000: "M"}
//	var numSlice []int
//
//	var newNumSlice []int
//	for num > 0 {
//		reverse := num % 10
//		num = num / 10
//		numSlice = append(numSlice, reverse)
//	}
//	fmt.Println(numSlice)
//	for i := 0; i < len(numSlice); i++ {
//		element := numSlice[i] * int(math.Pow(float64(10), float64(i)))
//		fmt.Println(element)
//		newNumSlice = append(newNumSlice, element)
//	}
//	s := GetString(numSlice, newNumSlice)
//	fmt.Println(newNumSlice)
//	return s
//}
//
//func GetString(numSlice []int, newNumSlice []int) string {
//	var s string
//	for i := len(newNumSlice) - 1; i >= 0; i-- {
//		switch {
//		case newNumSlice[i] < 4:
//			for j := 0; j < newNumSlice[i]; j++ {
//				s = s + "I"
//			}
//		case newNumSlice[i] == 4:
//			s = s + "IV"
//		case newNumSlice[i] >= 5 && newNumSlice[i] <= 8:
//			s = s + "V"
//			for j := 0; j < newNumSlice[i]-5; j++ {
//				s = s + "I"
//			}
//		case newNumSlice[i] == 9:
//			s = s + "IX"
//		case newNumSlice[i] >= 10 && newNumSlice[i] < 40:
//			for j := 0; j < numSlice[i]; j++ {
//				s = s + "X"
//			}
//		case newNumSlice[i] == 40:
//			s = s + "XL"
//		case newNumSlice[i] >= 50 && newNumSlice[i] < 90:
//			s = s + "L"
//			for j := 0; j < numSlice[i]-5; j++ {
//				s = s + "X"
//			}
//
//		}
//
//	}
//	return s
//}
//func main() {
//	var nums = []int{-1, 0, 1, -2, 0, 2, -1, -1, -5, 3, -10, 7, -8}
//	result := ThreeSum(nums)
//	fmt.Println(result)
//
//}
//func ThreeSum(nums []int) [][]int {
//	sort.Ints(nums)
//	fmt.Println(nums)
//	var result [][]int
//	for i := 0; i < len(nums)-2; i++ {
//		fmt.Println(nums[len(nums)-2])
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//		left := i + 1
//		right := len(nums) - 1
//		for left < right {
//			sum := nums[i] + nums[left] + nums[right]
//			if sum < 0 {
//				left++
//			} else if sum > 0 {
//				right--
//			} else {
//				result = append(result, []int{nums[i], nums[left], nums[right]})
//				for left < right && nums[left] == nums[left+1] {
//					left++
//				}
//				for left < right && nums[right] == nums[right-1] {
//					right--
//				}
//				left++
//				right--
//			}
//		}
//	}
//	return result
//}
//// Функция, принимающая callback и выполняющая операцию с числом
//func operation(x int, callback func(int)) {
//	result := x * 2
//	callback(result)
//}
//
//// Callback-функция, которая просто выводит результат
//func printResult(result int) {
//	fmt.Println(result)
//}
//
//// Callback-функция, которая выполняет дополнительную операцию с результатом
//func multiplyBy3(result int) {
//	finalResult := result * 3
//	fmt.Println(finalResult)
//}
//
//func main() {
//	operation(5, printResult) // Выведет 10
//	operation(5, multiplyBy3) // Выведет 30
//}
//type Person struct {
//	Name string
//	Age  int
//}
//type My []Person
//
//func main() {
//	var nums = []int{7, 4, 5, 6}
//	sort.Ints(nums)
//	fmt.Println(nums)
//
//	var str = []string{"d", "b", "a", "e"}
//	sort.Strings(str)
//	fmt.Println(str)
//
//	person := []Person{{"Alice", 78}, {"Alex", 65}, {"Kate", 90}, {"Dima", 16}}
//	sort.Slice(person, func(i, j int) bool {
//		return person[i].Age > person[j].Age
//
//	})
//	fmt.Println(person)
//	sort.Sort(My(person))
//	fmt.Println(person)
//}
//
//func (a My) Len() int           { return len(a) }
//func (a My) Less(i, j int) bool { return a[i].Age < a[j].Age }
//func (a My) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
//ВОККЕРЫ
//func worker(id int, tasks <-chan int, results chan<- int) {
//	for task := range tasks {
//		fmt.Println("Worker", id, "started task", task)
//		time.Sleep(200 * time.Millisecond) // Имитация выполнения задачи
//		fmt.Println("пустая стока")
//		fmt.Println("Worker", id, "finished task", task)
//
//		results <- task * 2
//	}
//}
//
//func main() {
//	taskCount := 10
//	workerCount := 4
//
//	tasks := make(chan int, taskCount)
//	results := make(chan int, taskCount)
//
//	for w := 1; w <= workerCount; w++ {
//		go worker(w, tasks, results)
//	}
//
//	for t := 1; t <= taskCount; t++ {
//		tasks <- t
//	}
//
//	close(tasks)
//
//	for r := 1; r <= taskCount; r++ {
//		result := <-results
//		fmt.Println("Result:", result)
//	}
//
//}
//Создайте две горутины, чтобы числа из одного канала
//читались по мере поступления, возводились в квадрат и результат записывался во второй канал.
//func main() {
//	myChanel := make(chan int)
//	var nums = []int{2, 4, 5, 6, 7}
//	var wg sync.WaitGroup
//	for _, value := range nums {
//		wg.Add(1)
//		go func(value int) {
//			defer wg.Done()
//			myChanel <- value
//		}(value)
//	}
//	go func() {
//		wg.Wait()
//		close(myChanel)
//	}()
//	second := secondChanel(myChanel)
//	for val := range second {
//		fmt.Println(val)
//	}
//}
//func secondChanel(myChanel chan int) chan int {
//	secondChan := make(chan int)
//	done := make(chan struct{})
//	go func() {
//		defer close(secondChan)
//		for {
//			select {
//			case v, ok := <-myChanel:
//				if !ok {
//					return
//				}
//				secondChan <- v * v
//			case <-done:
//				return
//
//			}
//		}
//
//	}()
//
//	return secondChan
//}

//func main() {
//	var nums = []int{3, 4, 5, 7, 8}
//	firstChanel := make(chan int, len(nums))
//	var wg sync.WaitGroup
//	for _, value := range nums {
//		wg.Add(1)
//		go func(value int) {
//			defer wg.Done()
//			firstChanel <- value
//		}(value)
//	}
//	go func() {
//		wg.Wait()
//		close(firstChanel)
//	}()
//	second := secondChanel(firstChanel)
//	for v := range second {
//		fmt.Println(v)
//	}
//}
//func secondChanel(first chan int) chan int {
//	second := make(chan int)
//	done := make(chan struct{})
//	go func() {
//		defer close(second)
//		for {
//			select {
//			case v, ok := <-first:
//				if !ok {
//					return
//				}
//				second <- v * v
//			case <-done:
//				return
//				//case <-time.After(100 * time.Millisecond):
//				//	return
//			}
//		}
//	}()
//	return second
//}
//func main() {
//	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8}
//	result := getCouple(nums)
//	fmt.Println("результат", result)
//}
//func getCouple(nums []int) int {
//	var key int
//	numsMap := make(map[int]int)
//	for _, num := range nums {
//		if value, ok := numsMap[num]; ok {
//			numsMap[num] = value + 1
//		} else {
//			numsMap[num] = 1
//		}
//	}
//	for key, value := range numsMap {
//		fmt.Println("Ключ", key)
//		fmt.Println("Значение", value)
//		if value <= 1 {
//			return key
//		}
//	}
//	return key
//}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8}
	result := getCouple(nums)
	fmt.Println("результат", result)
}

func getCouple(nums []int) int {
	myMap := make(map[int]int)
	chanelMap := make(chan map[int]int)
	var wg sync.WaitGroup
	var mut sync.Mutex

	for _, num := range nums {
		wg.Add(1)
		go func(num int) {
			mut.Lock()
			defer wg.Done()
			if value, ok := myMap[num]; ok {
				myMap[num] = value + 1
			} else {
				myMap[num] = 1
			}
			chanelMap <- myMap
			mut.Unlock()
		}(num)
	}

	go func() {
		wg.Wait()
		close(chanelMap)
	}()
	result := getNum(chanelMap)
	return result
}

func getNum(chanelMap chan map[int]int) int {

	for v := range chanelMap {

		for key, value := range v {
			if value == 1 {
				return key
			}
		}
	}
	return 0
}
