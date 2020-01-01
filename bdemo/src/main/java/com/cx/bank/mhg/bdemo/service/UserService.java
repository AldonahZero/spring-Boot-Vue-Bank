package com.cx.bank.mhg.bdemo.service;

import com.cx.bank.mhg.bdemo.domain.TUser;
import com.cx.bank.mhg.bdemo.mapper.TUserDAO;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

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
}
