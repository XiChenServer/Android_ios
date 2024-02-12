package com.example.mine;

import android.os.Bundle;
import android.util.Log;

import androidx.appcompat.app.AppCompatActivity;

public class app_mine_MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        try {
            super.onCreate(savedInstanceState);
            setContentView(R.layout.activity_app_mine_main);
        } catch (Exception e) {
            Log.d("checkIsThereProblem", e.toString());
        }
    }
}