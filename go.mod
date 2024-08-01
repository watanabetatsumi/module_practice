module github.com/watanabetatsumi/module_practice/main

// replace training/greeting => ./greeting // ローカルパスへ置き換えるためモジュール名から相対パスを replace で渡す。

go 1.22.0

// replace training/greeting/hello => ./greeting/hello

require github.com/watanabetatsumi/module_practice/greeting v0.0.0-20240801021044-83fe678747cc
