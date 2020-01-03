package com.cx.bank.mhg.bdemo.service;

import com.cx.bank.mhg.bdemo.domain.TLog;
import com.cx.bank.mhg.bdemo.domain.TUser;
import com.cx.bank.mhg.bdemo.mapper.TLogDAO;
import com.cx.bank.mhg.bdemo.mapper.TUserDAO;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Service;

@Service
public class UserService {
    @Autowired
    private TUserDAO userDao;
    @Autowired
    private TLogDAO logDAO;
    @Autowired
    RedisTemplate redisTemplate;

    /**
     * 查余额
     */
    public TUser inquiry(TUser tUser) {
//        TUser ansUser = null;
//        if(tUser!=null){
//            ansUser = (TUser)redisTemplate.get(tUser.getUserId()+"");
//        }
        return userDao.selectByPrimaryKey(tUser.getUserId());
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

    /**
     * 用户注册
     */
    public int doRegister(TUser user) {
        user.setUserFlag(1);
        user.setBalance(0.0);
        int a =userDao.insert(user);
        return a;
    }

    /**
     * 用户存款
     */
    public int doAddMoney(TUser user, Double addNumber) {
        TUser fastUser = userDao.selectByPrimaryKey(user.getUserId());
        if (addNumber < 0 || fastUser == null){
            return -1;
        }
        fastUser.setBalance(fastUser.getBalance()+addNumber);
        int a =userDao.updateByPrimaryKey(fastUser);
        return a;
    }

    /**
     * 用户取款
     */
    public int doSubMoney(TUser user, Double subNumber) {
        TUser fastUser = userDao.selectByPrimaryKey(user.getUserId());
        if (fastUser == null || fastUser.getBalance() < subNumber){
            return -1;
        }
        fastUser.setBalance(fastUser.getBalance() - subNumber);
        int a =userDao.updateByPrimaryKey(fastUser);
        return a;
    }

    /**
     * 用户转账
     */
    public int doTransfer(Integer from, Integer to, Double transferNumber) {
        TUser fastfrom = userDao.selectByPrimaryKey(from);
        TUser fastto = userDao.selectByPrimaryKey(to);
        if (fastfrom == null || fastto == null ){
            return -1;
        }
        int fr = doSubMoney(fastfrom,transferNumber);
        if(fr!=1){
            return -1;
        }
        fr = doAddMoney(fastto,transferNumber);
        if(fr!=1) {
            doAddMoney(fastfrom,transferNumber);
            return -1;
        }
        TLog tLogFrom = new TLog();
        TLog tLogTo = new TLog();
        tLogFrom.setLogAmount(transferNumber);
        tLogTo.setLogAmount(transferNumber);
        tLogFrom.setLogType("支出");
        tLogTo.setLogType("收入");
        tLogFrom.setUserid(from);
        tLogTo.setUserid(to);
        logDAO.insert(tLogFrom);
        logDAO.insert(tLogTo);
        return 1;
    }
}
