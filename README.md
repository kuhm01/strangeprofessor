# Professorlang
# 이상한 나라의 교수님
이상한 나라의 교수님. 헤으응. 저 학생을 대학원으로 데려가야겠다!</br>
고 생각한 교수님. 그러자 학생들이 이상한 말을 쓰시기 시작했습니다.</br>

---
## 자료구조
### 정적변수
교수님이 말씀하시길 학과에는 교수님이 총 7분이 계사다고 합니다.</br>
그래서 우리는 7명의 교수님들에게 자료를 저장할 수 있습니다.
```Go
교수님??? AAA 주세요,
교수님? ABA 내놔,
교수님?? BBAC 줘,
```
교수님 뒤의 '?'의 갯수로 몇번째 교수님에게 요청할 것인지 결정합니다.</br>
예를들어 교수님??? 이면 3번째 교수님을 의미합니다.</br>

당연하지만 우리 학과 교수님은 바쁘셔서 한 사람당 한 개의 자료만 저장할 수 있습니다.</br>

### 스택
우리 학과는 기본적으로 10명의 재학생을 수용할 수 있습니다.</br>
왜 이렇게 적냐면 정부에게 미움받아 예산을 못 받아서 그렇습니다.</br>
하지만 걱정마십시오! 우리가 여석신청을 하면 학과는 적자를 무릅쓰고 신청을 받아줍니다.</br>
```Go
여석신청 12 자리 //재학생 스택이 12칸 늘어서 총 22가 된다.
```

### 큐
우리 학과에는 전과생 리스트들이 있습니다.</br>
이 학생들은 다시 우리 학과에 오고싶어해서 선입선출의 구조를 가집니다!</br>
큐는 공간의 제약이 따로 없습니다.</br>

---
## 문법
교수님이 사용하시는 이상한 나라의 말들의 문법입니다.

### 문장들의 나열
언어 : 이상한 나라의 교수님은 문장들의 집합으로 구성됩니다.</br>
따라서 당신은 문장들을 구분해야합니다.</br>
문장의 구분은 ","을 사용합니다.
```Go
전공 ast, 여석신청 12 자리, 졸업,
```

또한 문장을 작성할 때에는 각 키워드는 빈 칸으로 구분됩니다.</br>
위와 같이 전공 명령문은 전공 `<identifier>`의 형식으로</br>
여석신청 명령문은 여석신청 `<int>` 자리의 형식으로 되어집니다.</br>
</br>

위의 대전제를 지키지않으면 교수님이 분노하셔서 F(fault)를 주실 것입니다.</br>
또한 마지막에는 항상 '졸업' 키워드가 있어야합니다.</br>
그래야 교수님이 이 학생은 졸업하고 싶어하는구나! 를 아시고</br>
프로그램을 종료할 것입니다.(위의 예시코드와 같음)</br>


### 전공
전공 키워드를 사용하면 우리는 Program Counter에 tag를 남길 수 있습니다.
```Go
전공 computer_architecture
```

식별자는 반드시 "_"과 영어 알파벳 소문자로만 구성되어야합니다.</br>
왜 그렇냐면 우리 교수님은 해외유학파라서 한글을 못 쓰시기때문입니다.</br>

### 재수강
재수강 키워드를 사용하면 Program Counter는 당신이 전공과목을 다시 수강하여</br>
높은 학점을 받고자하는 성실한 학생으로 판단해서</br>
현재 읽고있는 Program Index를 위의 전공 키워드를 사용해 tag한 곳으로 이동합니다!</br>
```Go
전공 ast
...
...
...
재수강 ast
```
위의 코드에서 당신은 ast tag를 계속해서 재수강하게 될 것입니다.</br>
눈치채셨겠지만 이상한 나라의 교수님에서 for-loop는 바로 위의 키워드를 활용하여</br>
사용할 수 있습니다.</br>
당연하지만 재수강할려고 하는 데, 해당 tag가 없으면 우리의 미숙한 Program Counter는</br>
해당 상황을 이해하지 못 하고 막강한 오류들을 분출할 것입니다.</br>

### 졸업
졸업하세요! 참고로 필자는 제 때에 졸업하지 못 했습니다. 너무 슬프네요.</br>
여러분들은 그렇게되지마시라고 졸업 키워드를 알려드립니다.</br>
이상한 나라의 교수님은 졸업 키워드를 감지하는 순간 종료됩니다.</br>
따라서 프로그램의 마지막에는 항상 졸업 키워드를 쓰십시오.!</br>
```Go
졸업, 
//또는
졸업

```
위의 코드가 알려주는 것은 우리의 parser는 바보라서 </br>
졸업 키워드 뒤에 문장을 구분짓는 "," 또는 enter가 없으면 막강한 오류를 내뱉으며 종료됩니다.</br>
parser를 바보로 만들고싶지않으면 졸업 키워드 뒤에는 "," 또는 enter를 항상 사용해야합니다.</br>

### 여석신청
재학생의 수는 한정되어 있습니다. 하지만 좀 더 많은 재학생이 필요하다면</br>
다음과 같은 형식으로 학과에서 여석을 받아낼 수 있습니다.</br>
```Go
여석신청 12 자리
```
위의 코드를 입력하면 우리의 학과는 재학생의 한계를 12만큼 늘려줄 것입니다.</br>

### 교수님?? 학점 주세요
우리는 교수님이라는 주어진 7개의 변수 영역에 값을 저장할 수 있습니다.</br>
방법은 다음과 같음.
```Go
교수님?? ABA 주세요,
교수님??? AF 줘,
교수님? AAAA 내놔,
```
교수님 키워드를 작성하고 그 뒤에 '?'를 작성합니다.</br>
'?'의 갯수는 몇번째 교수님에게 값을 저장할 지를 나타냅니다.</br>
즉, 교수님??? 이면 3번째 교수님에게 값을 저장합니다.</br>

뒤의 학점들의 나열은 곧 정수조합을 의미합니다.</br>
각 학점들은 고유한 정수값을 의미하는 데, 다음과 같습니다.</br>
+ A : 4
+ B : 3
+ C : 2
+ D : 1
+ F : -1

</br>
학점은 대문자로 작성해야 합니다. 소문자는 인식하지 못 해욤~</br>
학점나열들은 모조리 덧셈됩니다.</br>
즉, 다음과 같습니다.

```Go
ABB // 10
CF // 2
```

그 다음 학점달라는 말을 남겨야합니다.</br>
멘트가 3종류가 있는 데, 각각 다음의 의미를 가집니다.</br>

+ 주세요 : 1배율로 저장
+ 줘 : 2배율로 저장
+ 내놔 : 3배율로 저장

위 멘트를 작성해서 값을 저장할 때 몇배율로 저장할 건지 결정할 수 있습니다!</br>
교수님에게 값이 저장될 때는 새로운 값이 덮씌워지는 것이 아니라 기존 값이 더해지는</br>
방식으로 저장됩니다.</br>
즉, 기존에 6이라는 값이 저장되어 있었고 10을 저장하고자하면 기존값 + 신규값</br>
6 + 10 = 16이라는 값이 저장됩니다.</br>

### 교수님!! ㅁㅁ발표
우리는 위의 명령문으로 교수님에 저장한 값을 불러와 버퍼에 저장할 수 있습니다.</br>
다음과 같이 활용</br>
```Go
교수님? A 줘, // 첫번째 교수님에 8 저장
교수님! 성적발표, // 버퍼에 8을 character형태로 저장
교수님! 점수발표, // 버퍼에 8을 정수형태로 저장
```
위와 같이 '!'의 갯수는 몇번째 교수님인지를 나타냅니다.</br>
성적발표 키워드는 버퍼에 값을 저장할 때 문자 형태로 저장합니다.</br>
</br> 
정확하게는 문자형태로 출력하라는 명령을 함께 저장합니다.</br>
반대로 점수발표 키워드는 정수형태로 저장합니다.</br>
위에서 소개했듯, 교수님은 유학파이기때문에 한글을 못 씁니다.</br>
</br>
그래서 아스키 코드에 있는 알파벳들만 출력할 수 있습니다.</br>
한글 써보세요 허허허허. 멍청한 우리의 parser(학생1)나 evaluator(학생2)은 이를 알아듣지 못 할 것입니다.</br>
마치 우리들과 같습니다.</br>

### 공지
공지 명령문은 단독으로 사용합니다.
```Go
공지,
```
위 명령문을 사용하면 버퍼에 있던 값을 모조리 출력합니다.</br>
버퍼의 크기는 무한정하지않고 20이라는 크기제한이 있어서</br>
잘 사용하세요</br>

### 휴/보강
휴강과 보강 키워드가 추가되었습니다.</br>
휴강 키워드는 Program Counter의 Index를 1 증가시킵니다.</br>
즉, 다음 명령문을 한 번 뛰어넘습니다.</br>
</br>

반대로 보강 키워드는 Program Counter의 Index를 1 감소시킵니다.</br>
즉, 이 전의 명령문을 한 번 더 시행합니다.</br>
단, 보강 키워드는 해당 프로그램 내에서 1회성입니다.</br>
Program Counter는 내부적으로 보강 키워드가 있었던 Index를 기억한 후,</br>
해당 Index가 다시 오게 되었을 때 이를 생략합니다.</br>

### 신규임용
우리는 신규임용 키워드를 통해 교수님들에 저장된 값을 완전 초기화할 수 있습니다!</br>
신규임용 키워드를 사용하면 7개의 변수 공간을 0으로 초기화합니다.</br>
단, 아직 값이 할당되지않아 Null인 공간은 continue합니다.</br>

### 재학생 전입/출
재학생 전출, 재학생 전입 명령문을 통해서 우리는</br>
스택 자료구조와 큐 자료구조 간의 데이터를 옮길 수 있습니다!</br>
재학생 전출 명령문은 스택에서 큐로 데이터를 하나 옮깁니다.</br>
```Go
재학생 전출,
```
반면 재학생 전입 명령문은 큐에서 스택으로 데이터를 하나 옮깁니다.</br>
```Go
재학생 전입,
```

### 입학했습니다
교수님... 입학했습니다 명령문을 사용하면</br>
n번째 교수님에 저장되어 있던 값을 재학생 stack에 저장합니다.</br>
```Go
교수님... 입학했습니다, //3번째 교수님에 있는 값을 스택에 저장
```

### 수료했습니다
수료했습니다 교수님.. 명령문을 사용하면</br>
n번째 교수님에 재학생 stack에서 값을 하나 가져와 저장합니디.</br>
```Go
수료했습니다 교수님... //스택에서 값을 하나 가져와 3번째 교수님에 저장
```

### 평점제출
평점제출 키워드를 활용하면 교수님이 가지고 계신 값을</br>
이른바, miniBuffer에 저장할 수 있습니다.</br>
miniBuffer는 값을 저장할 수 있는 단 두개의 공간을 가진</br>
정적구조로 miniBuffer에 저장된 값으로 연산을 시행하게 됩니다.</br>
사용예시
```Go
교수님.. 평점제출 //두번째 교수님에 저장된 값을 miniBuffer에 저장
```
위의 입학과 문법이 비슷하므로 어려운 것은 없습니다.</br>

---
연산은 어떻게 해요?</br>
A. 아직 안 만들었습니다.

---

---
## 실행해보기
컴파일해서 하나의 실행파일을 만들어도 되고 go에서 바로 실행해도 무방합니다.</br>
```Go
go run main.go //in main directory
//or
go build
professorc
```

교수님은 다음의 실행 옵션을 제공합니다.</br>
+ -h, help : 도움말 출력
+ -c, start : 파일 읽어 인터프리트 시작

### 파일 읽기
이상한 나라의 교수님 언어는 .gyosoo 의 확장자를 가집니다.
예시</br>
+ main.gyosoo
+ hello.gyosoo

파일을 읽어 실행할려면 다음의 명령문을 입력하세요.</br>
```Go
go run main.go -c file directory/filename.gyosoo
//ex. go run main.go -c test/hello.gyosoo
//or
professorc -c file directory/filename.gyosoo
//ex. professorc -c test/hello.gyosoo
```

만일 올바르지 못 한 파일. 또는 이상한 입력 등이 있으면</br>
교수님은 인터프리터 불가 판정을 내려주실 것입니다.</br>

---
## TestCase
main directory 안, test 폴더에 이상한 나라의 교수님 언어가</br>
반드시 통과해야 할 표준 명령들이 있습니다.</br>
이상한 나라의 교수님 언어의 구현체는 이러한 표준 케이스를</br>
반드시 통과해야 합니다.</br>
+ hello.gyosoo : Hello World 출력
+ qstest.gyosoo : 큐와 스택 간 자료이동 시험 //큐와 스택에 공간이 없음을 알리는 오류문 출력 후, H72 출력

---
## 튜링 완전?
이상한 나라의 교수님은 튜링 완전한가요?</br>
네. 이제는 그렇다고 주장할 수 있게되었습니다.</br>
다음을 봐주십시오. 학생들이 새롭게 말하기 시작한 문법입니다.</br>

---
### 새로운 문법
### 교양, 교양필수, 교양선택, 교양끝
이상한 나라의 교수님 언어에 드디어 최신식? 조건분기문이</br>
등장했습니다.</br>
다음과 같이 사용합니다.
```Go
교양,
교양필수 .....,
교양선택 .....,
교양끝,
```
교양 키워드는 내부에 조건분기를 조정하는 변수를</br>
검토하여 분기여부를 결정합니다.</br>
조건분기 변수를 검토하여 참(true)일 경우, 교양필수.</br>
조건분기 변수를 검토하여 거짓(falase)이면, 교양선택.</br>
을 실행하게됩니다.</br>
.....은 당신이 입력한 다른 다양한 명령문들입니다.</br>

마지막은 항상 교양끝 키워드가 위치해야합니다.</br>
교양 명령문은 교양필수와 교양선택 키워드에 따라</br>
차례대로 명령문들을 읽어 wrapping하는 데,</br>
그 끝을 알리는 키워드가 교양끝 키워드입니다.</br>
물론 끝에 항상 ","을 입력하는 걸 잊지마세요!</br>

---
근데 그러면 조건분기문은 교양키워드들로 사용하는 건 알겠는 데,</br>
그 조건분기는 대체 어떻게 하나요?

---
서두르지마세요. 아래에서 그 방법을 소개합니다.</br>

---
### 조건 분기
이상한 나라의 교수님에서 조건 분기를 검토하는 두 변수는</br>
첫 번째 교수님과 두 번째 교수님을 기준으로 합니다.</br>
즉, 교수님. 이 교수님.. 보다 작으면 참(true)</br>
교수님. 이 교수님.. 보다 같거나 크면 거짓(false)</br>
입니다!</br>
아래에서 소개
```Go
교수님?? A 주세요,
교수님? C 주세요, 
///첫 번째 교수님은 4
///두 번째 교수님은 2
교양 //"첫 번째 교수님 < 두 번째 교수님" 이 거짓
/// 따라서 교양선택 문이 실행
교양필수
.....,
교양선택
.....,
교양끝
```
위와 같습니다.</br>
즉, 교양문을 호출하면 이상한 나라의 교수님 구현체는</br>
자동으로 내부에서 Professor[0]과 Professor[1]의 값을</br>
상호 비교하고 boolean 결과를 반환합니다.</br>
이를 바탕으로 우리의 ~~평가기가 실행됩니다.~~평가기는 아직 안 만들었습니다.</br>

---
### 왜 첫 번째, 두 번째 교수님인가요?
왜냐하면 첫 번째 교수님은 __학과장__</br>
두 번째 교수님은 __ABEEK 위원회장__</br>
에 재임 중이시기 때문입니다</br>
한 마디로 고생하시는 것이죠.</br>
