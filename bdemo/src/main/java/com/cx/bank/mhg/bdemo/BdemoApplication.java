package com.cx.bank.mhg.bdemo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class BdemoApplication {

	public static void main(String[] args) {
		SpringApplication.run(BdemoApplication.class, args);
//		ConfigurableApplicationContext context = SpringApplication.run(BdemoApplication.class, args);
//
//		KafkaSender sender = context.getBean(KafkaSender.class);
//
//		for (int i = 0; i < 3; i++) {
//			//调用消息发送类中的消息发送方法
//			sender.send();
//
//			try {
//				Thread.sleep(3000);
//			} catch (InterruptedException e) {
//				e.printStackTrace();
//			}
//		}
	}
}
