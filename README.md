#                         goModulePackageSetting
> golang的模块和包的使用设置
>
> >GoModule的简单介绍以及作用【个人理解】
> >
> >>1. **Go Moudule 即 Go 1.11** 之后官方支持的版本管理工具 mod
> >>
> >>2. 通过Go Mod 减轻了对GOPATH的依赖可以方便的安装从其他途径获取的包（一般是github的一些项目）
> >>
> >>3. 例如：
> >>
> >>   ![image-20210114161800727](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20210114161800727.png)
> >>
> >>4. 通过在main中使用go mod init main 会在当前文件夹下生成go.mod文件，里面记录了所引用的package
> >>
> >>5. 当通过go build -o 【当前目录】 【输出目录 】 windows下可以生成main.exe文件
> >>
> >>   linux,macos等同样会生成。
> >
> >>1. Go Package 使用 
> >>2. ![image-20210114162900117](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20210114162900117.png)
> >>3. 以上图为例：golang中的package 一般对应用 每一个目录名，同一个目录中的package名一样都是对于mypackage来说mypackage.go和mypackagetwo.go都是其中的*[portion](javascript:;)*都可以直接使用。

