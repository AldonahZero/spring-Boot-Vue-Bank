package com.cx.bank.mhg.bdemo.controller;

import com.cx.bank.mhg.bdemo.domain.TUser;
import com.cx.bank.mhg.bdemo.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/user")
public class UserController {
    @Autowired
    private UserService userService;

    @RequestMapping("/query")
    public TUser testQuery() {
        return userService.selectUserById(1);
    }


}