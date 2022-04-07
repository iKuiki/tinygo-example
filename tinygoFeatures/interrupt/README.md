# interrupt 中断

在单片机上，嵌入式应当是非常重要的一个基础功能，所以对于tinygo上interrupt的实现的探索，非常有意义。

在我刚开始使用tinygo的interrupt时，因为有着服务端go的开发经验，但是缺乏单片机的开发经验，所以踩了非常多的坑，甚至怀疑过是不是tinygo的interrupt在stm32下的实现有问题。

后来在github上翻到一个[别人写的demo](https://github.com/sago35/tinygo-examples/blob/main/pininterrupt/main.go)并且截取部分代码运行后，发现interrupt并无问题，还是我自己对单片机的不够熟悉踩到的坑。

这里用一个简单的开关按下-切换led灯状态的demo来演示interrupt的基本用法。
