package com.example.sellcowhourse;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;

import com.alibaba.android.arouter.facade.annotation.Route;

/**
 * @author winiymissl
 */
@Route(path = "/sellcowhourse/app_MainActivity")
public class app_MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_app_main);

    }
}