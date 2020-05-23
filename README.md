報錯是提示文字定時器引起，在組件unmounted之後再執行setstate會報溢出

sestate 是異步的，不能保證設置後即刻更新數據

setstate is async, can't guarantee update state data immediately

