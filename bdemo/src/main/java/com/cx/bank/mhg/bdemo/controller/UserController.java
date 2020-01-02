package com.cx.bank.mhg.bdemo.controller;

import ch.qos.logback.core.net.SyslogOutputStream;
import com.cx.bank.mhg.bdemo.domain.TUser;
import com.cx.bank.mhg.bdemo.exception.ResourceNotFoundException;
import com.cx.bank.mhg.bdemo.service.UserService;
import com.google.common.collect.ImmutableMap;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/user")
public class UserController {
    @Autowired
    private UserService userService;

    @PostMapping("/login")
    public ResponseEntity doLogin(@RequestBody TUser user) {
        TUser tUser = userService.doLogin(user);
        if(tUser == null){
            throw new ResourceNotFoundException(ImmutableMap.of("用户 id:", user.getUserId()));
        }
        return ResponseEntity.ok(tUser);
    }

    @PostMapping("/register")
    public ResponseEntity doRegister(@RequestBody TUser user) {
        int a = userService.doRegister(user);
        if(a != 1){
            throw new ResourceNotFoundException(ImmutableMap.of("用户 id:", user.getUserId()));
        }
        return ResponseEntity.ok(a);
    }

    @PutMapping("/deposit")
    public ResponseEntity deposit(@RequestBody TUser user) {
        double addNumber = user.getBalance();
        int a = userService.doAddMoney(user,addNumber);
        if(a != 1){
            throw new ResourceNotFoundException(ImmutableMap.of("存款为负数:", addNumber));
        }
        return ResponseEntity.ok(a);
    }

    @PutMapping("/withdrawals")
    public ResponseEntity withdrawals(@RequestBody TUser user) {
        System.out.println(user);
        double subNumber = user.getBalance();
        int a = userService.doSubMoney(user,subNumber);
        if(a != 1){
            throw new ResourceNotFoundException(ImmutableMap.of("取款超出余额:", subNumber));
        }
        return ResponseEntity.ok(a);
    }

    @PostMapping("/inquiry")
    public ResponseEntity inquiry(@RequestBody TUser user) {
        TUser tUser = userService.inquiry(user);
        if(tUser == null){
            throw new ResourceNotFoundException(ImmutableMap.of("用户 id:", user.getUserId()));
        }
        return ResponseEntity.ok(tUser);
    }

    @PostMapping("/transfer")
    public ResponseEntity transfer(@RequestBody Map<String, String> map) {
        System.out.println(map);
        Integer fromID = Integer.parseInt(map.get("fromId"));
        Integer toID = Integer.parseInt(map.get("toId"));
        Double transferNumber = Double.parseDouble(map.get("transferNumber"));
        System.out.println(fromID +" "+toID+" "+transferNumber);
        int a = userService.doTransfer(fromID,toID,transferNumber);
        if(a != 1){
            throw new ResourceNotFoundException(ImmutableMap.of("转账失败:", a));
        }
        return ResponseEntity.ok(a);
    }
}