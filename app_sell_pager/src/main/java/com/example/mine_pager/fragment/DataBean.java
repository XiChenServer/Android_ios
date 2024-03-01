package com.example.mine_pager.fragment;

/**
 * @Author winiymissl
 * @Date 2024-02-29 14:01
 * @Version 1.0
 */
public class DataBean {
    String url;

    public DataBean(String url) {
        this.url = url;
    }

    public String getUrl() {
        return url;
    }

    @Override
    public String toString() {
        return "DataBean{" +
                "url='" + url + '\'' +
                '}';
    }

    public void setUrl(String url) {
        this.url = url;
    }

}
