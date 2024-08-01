module training/main

replace training/greeting => ./greeting // ローカルパスへ置き換えるためモジュール名から相対パスを replace で渡す。

go 1.18

require training/greeting v0.0.0-00010101000000-000000000000

require training/greeting/hello v0.0.0-00010101000000-000000000000 // indirect

replace training/greeting/hello => ./greeting/hello
