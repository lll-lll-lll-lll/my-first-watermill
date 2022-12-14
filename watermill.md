# Watermillとは
- メッセージストリームを効率的に処理するためのライブラリ
- イベントドリブンなアプリケーションを作ることを目的にしている
- Pub/Subの実装一式が付属されている
- イベント駆動型アプリケーションの基本的な考え方は同じ
  1. メッセージの着信を待つ
  2. それに対してリアクションする
- PubSub ライブラリは複雑な機能を持っているが,WatermillではPublisherとSubscriberの２つのインターフェースを実装するだけ

## Wattermillのコアは`Message`
- http pkgでhttp.Requestが重要なように、どこでもMessage sturctは何かしら関わっている
- `Publisher`が発信する
- `Subscriber`が受信する
- メッセージの処理が完了したら、`Ack()`を送信する
- 処理が失敗したら`Nack()`を送信する
- `Acks()`と`Nacks()`は`Subscribers`に処理される
  - デフォでは`Subscribers`が`Ack`と`Nack`を待つ
- `NewMessage()`はペイロードとしてバイトのスライスを想定している
    - バイトにシリアライズできるなら、json, protobuf, Avro, gobなど使える


# Publisher
[公式Doc](https://watermill.io/docs/pub-sub/#publisher)
- ほとんどのpublisherはアトミックをサポートしていない。これはあるmessageがpublishに失敗した時に次のメッセージがpublishできないから
- 同期、非同期両方可能。実装に依存する
- publisherが非同期の場合`Close`は未送信のメッセージをフラッシュする必要がある<br>
- **publisherの閉じ忘れには注意する必要がある**
- そうしないとメッセージの一部が失われる可能性がある


# MessageのSubscribingについて
- トピック名を受け取り、受信したメッセージのチャネルを返す
- トピックが何を意味するかはPubSubの実装に依存する
- `Subscriber`の責務は`Ack`と`Nack`を処理すること
- メッセージのstorage/brokerへのAck/offsetはWatermillのメッセージからのAckの後に送らなければならない
- `Close`は全てのサブスクリプションとその出力チャネルを閉じ、必要な時にオフセットなどをフラッシュする