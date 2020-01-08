package com.cx.bank.mhg.bdemo.Utils;

import java.math.BigInteger;
import java.security.MessageDigest;

public class MD5Utils {
    
	public static String md5(String value) {
		byte[] secretBytes = null;
		try {
			secretBytes=MessageDigest.getInstance("md5").digest(value.getBytes());
		}catch (Exception e) {
			// TODO: handle exception
			e.printStackTrace();
		}
		String str=new BigInteger(1,secretBytes).toString(16);
		for (int i = 0; i < 32 - str.length(); i++) {
			str = "0" + str;
		}
		return str;
	}
	public static void main(String[] args) {
		System.out.println(md5("1234"));
	}
	
}
