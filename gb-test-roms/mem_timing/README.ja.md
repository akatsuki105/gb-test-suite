# GameBoyのCPUによるメモリアクセス時間のテスト

このROMは、スタックやプログラムカウンタへのアクセスを除く、CPU命令によるメモリの読み書きのタイミングをテストします。

これらのテストでは、正しい命令のタイミングと適切なタイマー操作（TAC、TIMA、TMA）が要求されます。

読み書きテストでは、失敗した命令を以下のようにリストアップしています。

```
[CB] opcode:tested-correct
```

read-modify-writeテストでは、失敗した命令を以下のようにリストアップします。

```
[CB] opcode:tested read/tested write-correct read/correct write
```

オペコードの後の値は、アクセスがどの命令サイクルで行われるかを示しており、1が最初の命令サイクルとなります。

他の問題でアクセスのサイクルが確定できなかった場合は0と表示されます。

読み出しか書き込みのどちらかで、両方ではない命令の場合、CPUは最後のサイクルでアクセスを行います。

読み取り、修正、書き戻しを行う命令では、CPUは最後の1個前のサイクルで読み取り、最後のサイクルで書き込みを行います。

## テスト内部の処理について

テストでは、タイマーが64サイクルごとにTIMAをインクリメントし、これに同期して可変量の遅延を行い、テスト対象の命令がタイマーにアクセスします。

遅延時間を1サイクル単位で変化させることで、命令によるメモリアクセスをTIMAのインクリメントの前後で行うようになっています。

その後、TIMAのレジスタと値を調べることで、どちらが発生したかを判断することができます。

## Multi-ROM

[cpu_instrs](../cpu_instrs/README.ja.md#multi-rom)参照

## エラーコード

[cpu_instrs](../cpu_instrs/README.ja.md#エラーコード)参照


## 画面出力

[cpu_instrs](../cpu_instrs/README.ja.md#画面出力)参照

## ソースコード

[cpu_instrs](../cpu_instrs/README.ja.md#ソースコード)参照

## Internal framework operation

[cpu_instrs](../cpu_instrs/README.ja.md#internal-framework-operation)参照

