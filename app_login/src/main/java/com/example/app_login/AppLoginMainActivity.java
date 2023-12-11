package com.example.app_login;

import androidx.appcompat.app.AppCompatActivity;
import androidx.fragment.app.Fragment;
import androidx.viewpager.widget.ViewPager;

import android.os.Bundle;
import android.widget.TableLayout;

import com.example.app_login.adapter.LoginFragmentAdapter;
import com.example.app_login.databinding.ActivityAppLoginMainBinding;
import com.google.android.material.tabs.TabItem;
import com.google.android.material.tabs.TabLayout;

import java.util.ArrayList;
import java.util.List;

/**
 * @author winiymissl
 */
public class AppLoginMainActivity extends AppCompatActivity {
    ActivityAppLoginMainBinding binding;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        binding = ActivityAppLoginMainBinding.inflate(getLayoutInflater());
        super.onCreate(savedInstanceState);
        setContentView(binding.getRoot());
        List<Fragment> fragmentList = new ArrayList<>();
        ViewPager viewPager = binding.vp;
        fragmentList.add(new SignInFragment("登录"));
        fragmentList.add(new RegisterFragment("注册"));
        LoginFragmentAdapter loginFragmentAdapter = new LoginFragmentAdapter(getSupportFragmentManager(), fragmentList);
        viewPager.setAdapter(loginFragmentAdapter);
        binding.tl.setupWithViewPager(viewPager);
        for (int i = 0; i < binding.tl.getTabCount(); i++) {
            TabLayout.Tab tabAt = binding.tl.getTabAt(i);
            if (fragmentList.get(i) instanceof RegisterFragment) {
                tabAt.setText("注册");
            }
            if (fragmentList.get(i) instanceof SignInFragment) {
                tabAt.setText("登录");
            }
        }
    }
}