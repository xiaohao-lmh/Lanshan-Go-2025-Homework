package main

import (
	"bufio"
	"fmt"
	"math/rand" //ai告诉可以用v2版本，不需要加seed,之后看看
	"os"
	"time"
)

type MentalState interface {
	Describe() string
	ReactToPressure(pressure int) MentalState //针对压力的反应
	GetStateName() string
} //接口

type Wakabamustusmi struct {
	Name      string
	Stress    int
	State     MentalState
	MaxStress int
}

type NormalState struct{} //正常睦

func (n NormalState) Describe() string {
	return "小睦在和月之森富婆一起种黄瓜"
}
func (n NormalState) ReactToPressure(pressure int) MentalState {
	if pressure >= 80 {
		return MortisState{}
	} else if pressure >= 60 {
		return Binglengmu{} //这个不需要了状态
	}
	return n
}

func (n NormalState) GetStateName() string {
	return "普通黄瓜人"
}

type Binglengmu struct{} //冰冷睦

func (b Binglengmu) Describe() string {
	descriptions := []string{
		"若叶睦呼吸急促",
		"若叶睦眼神迷离",
		"若叶睦说出'不会长久'",
	}
	return descriptions[rand.Intn(len(descriptions))]
}

func (b Binglengmu) ReactToPressure(pressure int) MentalState {
	if pressure >= 80 {
		return MortisState{}
	} else if pressure < 60 {
		return NormalState{}
	}
	return b
}
func (b Binglengmu) GetStateName() string {
	return "压力爆大"
}

type MortisState struct{} //莫提斯

func (m MortisState) Describe() string {
	descriptions := []string{
		"若叶睦吉他弹呲，精神崩溃，整个人摊在椅子上，彻底被Mortis所取代",
		"人偶的内心被替换为mortis",
	}
	return descriptions[rand.Intn(len(descriptions))]
}
func (m MortisState) ReactToPressure(pressure int) MentalState {
	if pressure < 60 {
		return NormalState{}
	} else if pressure >= 60 && pressure < 80 {
		return Binglengmu{}
	}
	return m
}
func (m MortisState) GetStateName() string {
	return "Mortis"
}
func NewWakaba() *Wakabamustusmi {
	return &Wakabamustusmi{
		Name:      "WakabaMutsumi",
		Stress:    0,
		State:     NormalState{},
		MaxStress: 100,
	}
}
func (w *Wakabamustusmi) ReceivePressure(pressure int) {
	fmt.Printf("若叶睦收到了%d点压力\n", pressure)
	w.Stress += pressure
	if w.Stress > w.MaxStress {
		w.Stress = w.MaxStress //防止压力值溢出（AI告诉的）
	}

	previousState := w.State.GetStateName()
	w.State = w.State.ReactToPressure(w.Stress)
	currentState := w.State.GetStateName()

	if previousState != currentState {
		fmt.Printf("人格切换：%s - %s\n", previousState, currentState)
	}
}

func (w *Wakabamustusmi) Rest() {
	reduction := 25 + rand.Intn(10)
	w.Stress -= reduction
	if w.Stress < 0 {
		w.Stress = 0
	}
	fmt.Printf("若叶睦和Soyo妈妈进行了sox，压力释放了%d，当前压力：%d\n", reduction, w.Stress)
	previousState := w.State.GetStateName()
	w.State = w.State.ReactToPressure(w.Stress)
	currentState := w.State.GetStateName()
	if previousState != currentState {
		fmt.Printf("人格切换: %s - %s\n", previousState, currentState)
	}
}

func (w *Wakabamustusmi) DisplayStatus() {
	fmt.Printf("---------当前状态---------\n")
	fmt.Printf("压力水平：%d/%d\n", w.Stress, w.MaxStress)
	fmt.Printf("心理状态:%s\n", w.State.GetStateName())
	fmt.Printf("状态描述:%s\n", w.State.Describe())
}

type StressSourse interface {
	GeneratePressure() int
	GetDescription() string
}

type StressFromNyamuchi struct{}

func (Nyam StressFromNyamuchi) GeneratePressure() int {
	return 20 + rand.Intn(5)
}

func (Nyam StressFromNyamuchi) GetDescription() string {
	return "喵梦摘掉睦的面具"
}

type StressFromSaki struct{}

func (Saki StressFromSaki) GeneratePressure() int {
	return 20 + rand.Intn(5)
}

func (Saki StressFromSaki) GetDescription() string {
	return "祥子：为什么不帮我说话，明明我只剩下mujica了！"
}

type StressFromBand struct{}

func (Band StressFromBand) GeneratePressure() int {
	return 30 + rand.Intn(5)
}
func (Band StressFromBand) GetDescription() string {
	return "车站前乐队濒临崩溃"
}
func waitforinput() {
	fmt.Printf("按回车键继续")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	Wakabamustusmi := NewWakaba()
	Wakabamustusmi.DisplayStatus()

	StressSourses := []StressSourse{
		StressFromSaki{},
		StressFromNyamuchi{},
		StressFromBand{},
	}
	fmt.Println("月光给予了人偶生命......")
	for {
		waitforinput()
		sourse := StressSourses[rand.Intn(len(StressSourses))]
		pressure := sourse.GeneratePressure()

		fmt.Printf("时间流逝，%s\n", sourse.GetDescription())
		Wakabamustusmi.ReceivePressure(pressure)
		Wakabamustusmi.DisplayStatus()

		if Wakabamustusmi.Stress >= Wakabamustusmi.MaxStress {
			fmt.Println("每次苏醒皆是重生，\n不再苏醒便是永恒的死亡,\n我愿为你编织摇篮，\n以天鹅绒缝制,\n灵柩般的摇篮曲,\n将你包裹，\n你甚至不会察觉,自己陷入了沉眠\n所以,已经没事了。")
			fmt.Println("晚安，祝你好梦")
			break
		}

		if rand.Intn(3) == 0 {
			Wakabamustusmi.Rest()
			Wakabamustusmi.DisplayStatus()
		}
		time.Sleep(1 * time.Second)
	}
}
