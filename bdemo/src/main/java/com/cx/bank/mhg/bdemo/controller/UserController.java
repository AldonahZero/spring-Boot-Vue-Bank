package com.cx.bank.mhg.bdemo.controller;

import ch.qos.logback.core.net.SyslogOutputStream;
import com.cx.bank.mhg.bdemo.domain.TUser;
import com.cx.bank.mhg.bdemo.exception.ResourceNotFoundException;
import com.cx.bank.mhg.bdemo.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

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
            Map error = new HashMap();
            error.put("用户 id:", user.getUserId());
            throw new ResourceNotFoundException(error);
        }
        return ResponseEntity.ok(tUser);
    }

    @PostMapping("/register")
    public ResponseEntity doRegister(@RequestBody TUser user) {
        int a = userService.doRegister(user);
        if(a != 1){
            Map error = new HashMap();
            error.put("用户 id:", user.getUserId());
            throw new ResourceNotFoundException(error);
        }
        return ResponseEntity.ok(a);
    }

}