package com.example.app_login.adapter;

import androidx.annotation.NonNull;
import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentManager;
import androidx.fragment.app.FragmentPagerAdapter;

import java.util.List;

/**
 * @Author winiymissl
 * @Date 2023-12-11 14:26
 * @Version 1.0
 */
public class LoginFragmentAdapter extends FragmentPagerAdapter {
    List<Fragment> list;

    public LoginFragmentAdapter(@NonNull FragmentManager fm, List<Fragment> list) {
        super(fm);
        this.list = list;
    }

    @NonNull
    @Override
    public Fragment getItem(int position) {
        return list.get(position);
    }

    @Override
    public int getCount() {
        return list.size();
    }
}
