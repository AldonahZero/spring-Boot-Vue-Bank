package com.cx.bank.mhg.bdemo.mapper;

import com.cx.bank.mhg.bdemo.domain.TLog;
import org.springframework.stereotype.Repository;

/**
 * TLogDAO继承基类
 */
@Repository
public interface TLogDAO extends MyBatisBaseDao<TLog, Integer> {
}