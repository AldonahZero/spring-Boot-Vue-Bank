package com.cx.bank.mhg.bdemo.domain;

import lombok.Data;

import java.util.Date;

@Data
public class Message {
    private Long id;    //id

    private String msg; //消息

    private Date sendTime;  //时间戳

}