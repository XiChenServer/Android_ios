package com.example.sellcowhourse;

import android.app.Application;

import com.alibaba.android.arouter.launcher.ARouter;
import com.example.common.room.AppDatabase;
import com.tencent.mmkv.BuildConfig;
import com.tencent.mmkv.MMKV;

/**
 * @Author winiymissl
 * @Date 2023-12-12 21:19
 * @Version 1.0
 */
public class MyApplication extends Application {
    @Override
    public void onCreate() {
        super.onCreate();
        MMKV.initialize(this);
        if (BuildConfig.DEBUG) {
            ARouter.openLog();
            ARouter.openDebug();
        }
        ARouter.openLog();
        ARouter.init(this);
        //数据库的创建
        AppDatabase.getInstance(this);
    }
}
