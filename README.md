# es_notice
本プロジェクトは、以下の外部リソースの情報を定期的に監視し、状況に応じてメール発信する仕組みです。

　　　"http://wdc.nict.go.jp/IONO/fxEs/latest-fxEs.html" 

上記の情報を元に、電離層の状態でスポラディックE層の発生を通知します。


開発言語：golang <br />
サーバ環境：AWS EC2 (ubuntu)<br />
AWS :  SES(メール送受信サービス）<br />
<br />
(セットアップ方法)<br />
１．お手持ちのlinuxにgolangをインストール（方法は別サイト参照）<br />
２．必要なgoライブラリをインストール<br />
<space>    golang.org/x/net/html<br />
<space>    github.com/aws/aws-sdk-go/aws<br />
<space>    github.com/aws/aws-sdk-go/aws/session<br />
<space>    github.com/aws/aws-sdk-go/service/ses<br />
３．awsのSES(メールサービス)を使えるようにする<br />
<space>　　　このソースでは、aws実行ユーザのcredential ファイル、を実行ユーザのホームディレクトリの配下の.awsに保存<br />
<space>　　　（例：　/home/ubuntu/.aws 　)　　（例：aws実行ユーザ　：yuichi01 )<br />
４．実行プログラムの生成　<br />
<space>　　es_checkweb.goを置いたディレクトリで以下を実行し、実行ファイル（es_checkweb）を生成する<br />
<space>    go build -o es_checkweb -i<br />
５．cronへの登録（方法は別サイト参照）　<br />
<space>       (crotab例：  10,25,40,55 6-19 * * * ubuntu cd (program's directory) ; ./es_checkweb <br />

<br />
Tsukiyama-do

---------------------------

The project contains modules that monitor the following homepage of other institutions regularly, and raises a alert message of email in certain conditions.

　　　"http://wdc.nict.go.jp/IONO/fxEs/latest-fxEs.html" 

It notifies sporadic E layer of the ionosphere according to the homepage.


Language：golang <br />
Server equipments：AWS EC2 (ubuntu)<br />
AWS :  SES(e-mail system）<br />

(How to setup)<br />
1. Install golang to your linux PC or your aws EC2  (See other sites in details)<br />
2. Install the below golang libraries<br />
<space>    golang.org/x/net/html<br />
<space>    github.com/aws/aws-sdk-go/aws<br />
<space>    github.com/aws/aws-sdk-go/aws/session<br />
<space>    github.com/aws/aws-sdk-go/service/ses<br />
3. Set up aws' SES(email transfer service)<br />
<space>In this code, credential file and config file of the aws user are placed under .aws directory under the home directory of the execution linux user.<br />
<space>　　　（example：　/home/ubuntu/.aws 　)　　（example：aws user ：yuichi01 )<br />
4. Create execution program from golang source code　<br />
<space>　　Run the following command under the directory of the file of es_checkweb.go to generateo es_checkweb module.<br />
<space>    go build -o es_checkweb -i<br />
5. Register the program to cron. (See other sites about cron)　<br />
<space>       (crotab sample：  10,25,40,55 6-19 * * * ubuntu cd (program's directory) ; ./es_checkweb <br />
<space>    
<br />

Tstukiyama-do
