package com.example.sellcowhourse;

import android.os.Bundle;
import android.view.MenuItem;

import androidx.annotation.NonNull;
import androidx.appcompat.app.AppCompatActivity;
import androidx.viewpager.widget.ViewPager;

import com.alibaba.android.arouter.facade.annotation.Route;
import com.example.app_login.adapter.LoginFragmentAdapter;
import com.example.mine.fragment.app_mine_MainFragment;
import com.example.mine_pager.fragment.BlankFragment;
import com.example.sellcowhourse.databinding.ActivityAppMainBinding;
import com.google.android.material.bottomnavigation.BottomNavigationView;

import java.util.ArrayList;
import java.util.List;

/**
 * @author winiymissl
 */
@Route(path = "/sellcowhourse/app_MainActivity")
public class app_MainActivity extends AppCompatActivity {
    ActivityAppMainBinding binding;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        binding = ActivityAppMainBinding.inflate(getLayoutInflater());
        super.onCreate(savedInstanceState);
        setContentView(binding.getRoot());
//        查看Drawable的所有子类
        List list = new ArrayList();
        list.add(new BlankFragment());
        list.add(new app_mine_MainFragment());
        LoginFragmentAdapter loginFragmentAdapter = new LoginFragmentAdapter(getSupportFragmentManager(), list);
        binding.mainVp.setAdapter(loginFragmentAdapter);
        binding.mainVp.addOnPageChangeListener(new ViewPager.OnPageChangeListener() {
            @Override
            public void onPageScrolled(int position, float positionOffset, int positionOffsetPixels) {

            }

            @Override
            public void onPageSelected(int position) {
                //根据碎片添加的顺序
                if (position == 0) {
                    binding.navView.setSelectedItemId(R.id.navi_search);
                } else if (position == 1) {
                    binding.navView.setSelectedItemId(R.id.navi_you);
                }
            }

            @Override
            public void onPageScrollStateChanged(int state) {

            }
        });
        binding.navView.setOnNavigationItemSelectedListener(new BottomNavigationView.OnNavigationItemSelectedListener() {
            @Override
            public boolean onNavigationItemSelected(@NonNull MenuItem item) {
                if (item.getItemId() == R.id.navi_you) {
                    binding.mainVp.setCurrentItem(1);
                    return true;
                } else if (item.getItemId() == R.id.navi_goods) {
//                    binding.mainVp.setCurrentItem();
                    return true;
                } else if (item.getItemId() == R.id.navi_search) {
                    binding.mainVp.setCurrentItem(0);
                    return true;
                } else if (item.getItemId() == R.id.navi_mess) {
//                    binding.mainVp.setCurrentItem();
                    return true;
                }
                return false;
            }
        });
    }
}