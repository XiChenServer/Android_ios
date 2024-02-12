package com.example.mine_pager;

import android.os.Bundle;
import android.util.Log;

import androidx.appcompat.app.AppCompatActivity;

public class app_sell_pager_MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        try {
            super.onCreate(savedInstanceState);
            setContentView(R.layout.activity_app_sell_pager_main);
        } catch (Exception e) {
            Log.d("ThereIsProblem", e.toString());
        }
    }
}