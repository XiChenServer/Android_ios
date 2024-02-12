package com.example.sellcowhourse;

import android.content.Intent;
import android.os.Bundle;
import android.os.Handler;
import android.util.Log;
import android.view.View;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;

import com.alibaba.android.arouter.facade.annotation.Route;
import com.alibaba.android.arouter.launcher.ARouter;
import com.google.android.material.chip.Chip;
import com.tencent.mmkv.MMKV;

@Route(path = "/sellcowhourse/SplashActivity")
public class SplashActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_splash);
        Chip chip = findViewById(R.id.chip);
        Handler handler = new Handler();
//        delAll();
        handler.postDelayed(new Runnable() {
            @Override
            public void run() {
                if (isLogin()) {
                    ARouter.getInstance().build("/sellcowhourse/app_MainActivity").withFlags(Intent.FLAG_ACTIVITY_NEW_TASK | Intent.FLAG_ACTIVITY_CLEAR_TASK).navigation();
                    overridePendingTransition(com.example.common.R.anim.set_in, com.example.common.R.anim.set_out);
                } else {
                    ARouter.getInstance().build("/app_login/AppLoginMainActivity").withFlags(Intent.FLAG_ACTIVITY_NEW_TASK | Intent.FLAG_ACTIVITY_CLEAR_TASK).navigation();
                    overridePendingTransition(com.example.common.R.anim.set_in, com.example.common.R.anim.set_out);
                }
            }
        }, 3000);

        chip.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                handler.removeCallbacksAndMessages(null);
                if (isLogin()) {
                    Toast.makeText(SplashActivity.this, "点击", Toast.LENGTH_SHORT).show();
                    try {
                        ARouter.getInstance().build("/app_login/AppLoginMainActivity").withFlags(Intent.FLAG_ACTIVITY_NEW_TASK | Intent.FLAG_ACTIVITY_CLEAR_TASK).navigation();
                        overridePendingTransition(com.example.common.R.anim.set_in, com.example.common.R.anim.set_out);
//                        Log.d("ThereIsProblem", "世界你好");
                    } catch (Exception e) {
                        Log.d("ThereIsProblem", e.toString());
                    }
                } else {
                    try {

                        ARouter.getInstance().build("/sellcowhourse/app_MainActivity").withFlags(Intent.FLAG_ACTIVITY_NEW_TASK | Intent.FLAG_ACTIVITY_CLEAR_TASK).navigation();
                        overridePendingTransition(com.example.common.R.anim.set_in, com.example.common.R.anim.set_out);
                    } catch (Exception e) {
                        Log.d("ThereIsProblem", e.toString());
                    }
                }
            }
        });
    }

    boolean isLogin() {
        //通过修改这里，进行调试和测试
        return MMKV.defaultMMKV().getString("token", null) != null;
//        Activity
//        startActivity();
    }

    void delAll() {
        MMKV.defaultMMKV().remove("token");
    }
}

