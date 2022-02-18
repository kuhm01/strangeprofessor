# TEST CASE
본 문서는 이상한 나라의 교수님 언어가</br>
통과해야 할 표준 케이스들의 모음입니다.</br>

1. __hello.gyosoo__
2. __illegal.gyosoo__
3. __qstest.gyosoo__


## hello.gyosoo
<details>
<summary>자세히 보기</summary>
<div markdown="1">
  
### Hello World
Hello World를 출력하는 기본 코드이다.
```Go
교수님??? AAAAAA 내놔,
교수님!!! 성적발표,
교수님??? AA 내놔,
교수님??? A 주세요,
교수님??? D 주세요,
교수님!!! 성적발표,
교수님??? AB 주세요,
교수님!!! 성적발표,
교수님!!! 성적발표,
교수님??? B 주세요,
교수님!!! 성적발표,
교수님? AAAA 줘,
교수님! 성적발표,
교수님?? AB 주세요,
교수님?? AAAAAAAAAA 줘,
교수님!! 성적발표,
교수님!!! 성적발표,
교수님??? B 주세요,
교수님!!! 성적발표,
교수님??? FFF 줘,
교수님!!! 성적발표,
교수님??? FFFF 줘,
교수님!!! 성적발표,
교수님???? AAC 주세요,
교수님!!!! 성적발표,
공지,
졸업,
```
  
  예상 결과 :</br>
  Hello World
</div>

</details>

## illegal.gyosoo
<details>
<summary>자세히 보기</summary>
<div markdown="1">
  
### illegal test
  illegal token을 검출하여 오류를 발생시키고</br>
  종료하는 코드이다.
```Go
몰루,
전공 교수,
졸업,
```
  
  예상 결과 :</br>
  ```
  Interpreter Error. don't ran Interpreting
  parser Error:
  몰루 is ILLEGAL
  expected next token to be IDENT, got=ILLEGAL instead
  교수 is ILLEGAL
  ```
</div>

</details>

## qstest.gyosoo
<details>
<summary>자세히 보기</summary>
<div markdown="1">
  
### Queue - Stack test
  Queue와 Stack간, Data 통신을 test하는 코드이다.
```Go
교수님? AAAAAAAAA 줘,
교수님! 성적발표,
재학생 전입,
재학생 전출,
교수님. 입학했습니다,
재학생 전출,
재학생 전입,
수료했습니다 교수님..,
교수님!! 점수발표,
공지,
졸업
  
```
  
  예상 결과 :</br>
  ```
  Queue is Empty
  There is No changed student
  Stack is Empty
  Student is not exist       
  H72
  ```
  처음 전입, 전출 명령에서는 Queue와 Stack에</br>
  저장 된 Data가 없으므로 이에 관한 오류가 출력된다.</br>
  그 뒤에는 Data가 있게되었으므로 과정을 거쳐 변수를 출력</br>
  ~~ is Empty는 연산과정에서 출력하는 오류이고</br>
  There is No changed student와 Student is not exist는</br>
  연산에서 오류를 전달받은 Queue와 Stack이 출력하는 오류이다.</br>
</div>

</details>
