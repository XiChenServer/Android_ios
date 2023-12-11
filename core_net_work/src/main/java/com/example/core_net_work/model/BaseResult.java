package com.example.core_net_work.model;

/**
 * @Author winiymissl
 * @Date 2023-12-11 13:17
 * @Version 1.0
 */
public class BaseResult {
    int code;
    String msg;

    public boolean isSucceed() {
        return code == ResultCodeType.OK;
    }

    public int getCode() {
        return code;
    }

    public void setCode(int code) {
        this.code = code;
    }

    public String getMsg() {
        return msg;
    }

    public void setMsg(String msg) {
        this.msg = msg;
    }
}
