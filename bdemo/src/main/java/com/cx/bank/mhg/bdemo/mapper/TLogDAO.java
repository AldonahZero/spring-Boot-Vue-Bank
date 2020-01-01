package com.cx.bank.mhg.bdemo.mapper;

import com.cx.bank.mhg.bdemo.domain.TLog;
import org.apache.ibatis.annotations.Mapper;
import org.springframework.stereotype.Repository;

/**
 * TLogDAO继承基类
 */
@Repository
@Mapper
public interface TLogDAO extends MyBatisBaseDao<TLog, Integer> {

    //转账
    int insert(TLog tLog);
}