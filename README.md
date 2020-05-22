# es_notice
本プロジェクトは、以下の外部リソースの情報を定期的に監視し、状況に応じてメール発信する仕組みです。

　　　"http://wdc.nict.go.jp/IONO/fxEs/latest-fxEs.html" 

上記の情報を元に、電離層の状態でスポラディックE層の発生を通知します。


開発言語：golang 
サーバ環境：AWS EC2 (ubuntu)
AWS :  SES(メール送受信サービス）

(セットアップ方法)
１．お手持ちのlinuxにgolangをインストール（方法は別サイト参照）
２．必要なgoライブラリをインストール
    golang.org/x/net/html
    github.com/aws/aws-sdk-go/aws
    github.com/aws/aws-sdk-go/aws/session
    github.com/aws/aws-sdk-go/service/ses
３．awsのSES(メールサービス)を使えるようにする
　　　このソースでは、aws実行ユーザのcredential ファイル、を実行ユーザのホームディレクトリの配下の.awsに保存
　　　（例：　/home/ubuntu/.aws 　)　　（例：aws実行ユーザ　：yuichi01 )
４．実行プログラムの生成　
　　es_checkweb.goを置いたディレクトリで以下を実行し、実行ファイル（es_checkweb）を生成する
    go build -o es_checkweb -i
５．実行シェル（es_check.sh）をcronへの登録（方法は別サイト参照）　

Tsukiyama-do

---------------------------

The project contains modules that monitor the following homepage of other institutions regularly, and raises a alert message of email in certain conditions.

　　　"http://wdc.nict.go.jp/IONO/fxEs/latest-fxEs.html" 

It notifies sporadic E layer of the ionosphere according to the homepage.


Language：golang 
Server equipments：AWS EC2 (ubuntu)
AWS :  SES(e-mail system）

Tstukiyama-do
