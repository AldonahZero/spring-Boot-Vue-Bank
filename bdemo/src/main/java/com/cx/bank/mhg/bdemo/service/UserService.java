package com.cx.bank.mhg.bdemo.service;

import com.cx.bank.mhg.bdemo.domain.TUser;
import com.cx.bank.mhg.bdemo.mapper.TUserDAO;
import com.sun.xml.internal.ws.util.Pool;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class UserService {
    @Autowired
    private TUserDAO userDao;

    /**
     * 根据名字查找用户
     */
    public TUser selectUserById(int id) {

        return userDao.selectByPrimaryKey(id);
    }

    /**
     * 用户登录
     */
    public TUser doLogin(TUser user) {
        TUser factUser =  userDao.selectByPrimaryKey(user.getUserId());
        if(factUser!=null){
            if(factUser.getUserPassword().equals(user.getUserPassword())){
                //登录成功， 会话缓存
                return factUser;
            }
        }
        // 登录失败
        return null;
    }
}
