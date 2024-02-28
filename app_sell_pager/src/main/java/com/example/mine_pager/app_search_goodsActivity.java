package com.example.mine_pager;

import android.content.Context;
import android.net.ConnectivityManager;
import android.net.NetworkInfo;
import android.os.Bundle;
import android.view.View;
import android.view.inputmethod.InputMethodManager;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.SearchView;

import com.alibaba.android.arouter.facade.annotation.Route;
import com.example.mine_pager.databinding.ActivitySearchGoodsBinding;

@Route(path = "/mine_pager/app_search_goodsActivity")
public class app_search_goodsActivity extends AppCompatActivity {
    ActivitySearchGoodsBinding binding;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        binding = ActivitySearchGoodsBinding.inflate(getLayoutInflater());
        setContentView(binding.getRoot());
        //默认展开
        binding.searchSearchView.setIconified(false);
        //变成两列
        //返回的集合，直接展示在页面，搜索一次，使用同一个适配器

        binding.searchSearchView.setOnQueryTextListener(new SearchView.OnQueryTextListener() {
            @Override
            public boolean onQueryTextSubmit(String query) {
                //首先对网络进行检测
                if (isNetworkConnected(app_search_goodsActivity.this)) {
                    Toast.makeText(app_search_goodsActivity.this, "网络连接良好", Toast.LENGTH_SHORT).show();
//                    List<Goods> list = new ArrayList<>();
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 40));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 40));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    list.add(new Goods(R.drawable.orenge, "橙子", "这是来自和克拉玛州的橙子,美味甘甜，真的好吃，比妈妈做的饺子还要好吃，买了不亏", 3.9f, 50));
//                    //这是来自服务器的list
//                    RecyclerViewAdapter recyclerViewAdapter = new RecyclerViewAdapter(list);
//                    GridLayoutManager gridLayoutManager = new GridLayoutManager(app_search_goodsActivity.this, 2);
//                    binding.searchRecyclerView.setAdapter(recyclerViewAdapter);
//                    binding.searchRecyclerView.setLayoutManager(gridLayoutManager);
                } else {
                    //进行文字显示
                    binding.searchTextView.setVisibility(View.VISIBLE);
                    Toast.makeText(app_search_goodsActivity.this, "网络走丢了", Toast.LENGTH_SHORT).show();
                }
                //取消聚焦，会自动关闭软键盘
                binding.searchSearchView.clearFocus();
                return true;
            }

            @Override
            public boolean onQueryTextChange(String newText) {
                return false;
            }
        });
    }

    private void shutSoftBroad() {
        InputMethodManager imm = (InputMethodManager) getSystemService(Context.INPUT_METHOD_SERVICE);
        imm.hideSoftInputFromWindow(getCurrentFocus().getWindowToken(), 0);
    }

    //检测是否存在网络
    private boolean isNetworkConnected(Context mContext) {
        if (mContext != null) {
            ConnectivityManager connectivityManager = (ConnectivityManager) mContext.getSystemService(Context.CONNECTIVITY_SERVICE);
            NetworkInfo networkInfo = connectivityManager.getActiveNetworkInfo();
            if (networkInfo != null) {
                boolean connected = networkInfo.isConnected();
                if (connected) {
                    if (networkInfo.getState() == NetworkInfo.State.CONNECTED) {
                        return true;
                    } else {
                        return false;
                    }
                }
            }
        }
        return false;
    }

    @Override
    public void onBackPressed() {
        super.onBackPressed();
        overridePendingTransition(com.example.common.R.anim.set_in, com.example.common.R.anim.set_out);
    }
}