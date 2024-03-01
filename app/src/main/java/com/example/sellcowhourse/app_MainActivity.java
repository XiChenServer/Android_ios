package com.example.sellcowhourse;

import android.os.Bundle;
import android.util.Log;
import android.view.MenuItem;

import androidx.annotation.NonNull;
import androidx.appcompat.app.AppCompatActivity;
import androidx.viewpager.widget.ViewPager;

import com.alibaba.android.arouter.facade.annotation.Route;
import com.example.app_login.adapter.LoginFragmentAdapter;
import com.example.mine.fragment.app_mine_MainFragment;
import com.example.mine_pager.fragment.BlankFragment;
import com.example.sellcowhourse.databinding.ActivityAppMainBinding;
import com.example.sellcowhourse.message.MessageFragment;
import com.example.sellcowhourse.shoppingcar.ShoppingCarFragment;
import com.google.android.material.bottomnavigation.BottomNavigationView;
import com.tencent.mmkv.MMKV;

import java.util.ArrayList;
import java.util.List;

/**
 * @author winiymissl
 */
@Route(path = "/sellcowhourse/app_MainActivity")
public class app_MainActivity extends AppCompatActivity {
    ActivityAppMainBinding binding;
    private static final int RC_CAMERA_PERMISSION = 123;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        binding = ActivityAppMainBinding.inflate(getLayoutInflater());
        super.onCreate(savedInstanceState);
        setContentView(binding.getRoot());

//        if (!EasyPermissions.hasPermissions(this, perms)) {
//            EasyPermissions.requestPermissions(this, "需要相机权限来拍摄照片", RC_CAMERA_PERMISSION, perms);
//        } else {
        // 已经有相机权限
        // 在这里执行您的逻辑
//        }
//        查看Drawable的所有子类
        List list = new ArrayList();
        list.add(new BlankFragment());
        list.add(new ShoppingCarFragment());
        list.add(new MessageFragment());
        list.add(new app_mine_MainFragment());
        LoginFragmentAdapter loginFragmentAdapter = new LoginFragmentAdapter(getSupportFragmentManager(), list);

        binding.mainVp.setAdapter(loginFragmentAdapter);
        //提前加载好四个页面
        binding.mainVp.setOffscreenPageLimit(4);
        binding.mainVp.addOnPageChangeListener(new ViewPager.OnPageChangeListener() {
            @Override
            public void onPageScrolled(int position, float positionOffset, int positionOffsetPixels) {

            }

            @Override
            public void onPageSelected(int position) {
                //根据碎片添加的顺序
                if (position == 0) {
                    binding.navView.setSelectedItemId(R.id.navi_search);
                } else if (position == 3) {
                    binding.navView.setSelectedItemId(R.id.navi_you);
                } else if (position == 1) {
                    binding.navView.setSelectedItemId(R.id.navi_goods);
                } else if (position == 2) {
                    binding.navView.setSelectedItemId(R.id.navi_mess);
                }
            }

            @Override
            public void onPageScrollStateChanged(int state) {

            }
        });
        Log.d("这是我的token", MMKV.defaultMMKV().getString("token",null));
        binding.navView.setOnNavigationItemSelectedListener(new BottomNavigationView.OnNavigationItemSelectedListener() {
            @Override
            public boolean onNavigationItemSelected(@NonNull MenuItem item) {
                if (item.getItemId() == R.id.navi_you) {
                    binding.mainVp.setCurrentItem(3, true);
                    return true;
                } else if (item.getItemId() == R.id.navi_goods) {
                    binding.mainVp.setCurrentItem(1, true);
                    return true;
                } else if (item.getItemId() == R.id.navi_search) {
                    binding.mainVp.setCurrentItem(0, true);
                    return true;
                } else if (item.getItemId() == R.id.navi_mess) {
                    binding.mainVp.setCurrentItem(2, true);
                    return true;
                }
                return false;
            }
        });
    }
}