package com.example.mine_pager;

import android.content.Context;
import android.net.ConnectivityManager;
import android.net.NetworkInfo;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.view.inputmethod.InputMethodManager;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.SearchView;
import androidx.fragment.app.FragmentManager;
import androidx.fragment.app.FragmentTransaction;
import androidx.recyclerview.widget.GridLayoutManager;

import com.alibaba.android.arouter.facade.annotation.Route;
import com.example.common.room.entitues.ProductSimple;
import com.example.common.room.recyclerviewitem.MyRecyclerViewItemClick;
import com.example.core_net_work.MyRetrofit;
import com.example.core_net_work.model.goods.ProductSimpleInfoResult;
import com.example.core_net_work.model.goods.SearchResult;
import com.example.mine_pager.adapter.RecyclerViewAdapter;
import com.example.mine_pager.databinding.ActivitySearchGoodsBinding;
import com.example.mine_pager.fragment.DetailFragment;

import java.util.ArrayList;
import java.util.List;
import java.util.function.Consumer;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

@Route(path = "/mine_pager/app_search_goodsActivity")
public class app_search_goodsActivity extends AppCompatActivity {
    ActivitySearchGoodsBinding binding;
    List<ProductSimple> list = new ArrayList<>();

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        binding = ActivitySearchGoodsBinding.inflate(getLayoutInflater());
        setContentView(binding.getRoot());
        //默认展开
        binding.searchSearchView.setIconified(false);
        //返回的集合，直接展示在页面，搜索一次，使用同一个适配器
        RecyclerViewAdapter recyclerViewAdapter = new RecyclerViewAdapter(list, app_search_goodsActivity.this);
        GridLayoutManager gridLayoutManager = new GridLayoutManager(app_search_goodsActivity.this, 2);
        binding.searchRecyclerView.setLayoutManager(gridLayoutManager);
        binding.searchRecyclerView.setAdapter(recyclerViewAdapter);
        binding.searchRecyclerView.addOnItemTouchListener(new MyRecyclerViewItemClick(this, new MyRecyclerViewItemClick.OnItemClickListener() {
            @Override
            public void onItemClick(View view, int position) {
                FragmentManager fragmentManager = getSupportFragmentManager();
                FragmentTransaction fragmentTransaction = fragmentManager.beginTransaction();
                DetailFragment fragment = new DetailFragment();
                Bundle bundle = new Bundle();
                bundle.putString("id", list.get(position).getId());
                Log.d("这有一个问题", list.get(position).getId());
                fragment.setArguments(bundle);
                fragmentTransaction.setCustomAnimations(com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right, com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right);
                fragmentTransaction.add(R.id.constraint_layout_search, fragment);
                fragmentTransaction.addToBackStack(null);
                fragmentTransaction.commit();
            }
        }));
        binding.searchSearchView.setOnQueryTextListener(new SearchView.OnQueryTextListener() {
            @Override
            public boolean onQueryTextSubmit(String query) {
                //首先对网络进行检测
                if (isNetworkConnected(app_search_goodsActivity.this)) {
                    Toast.makeText(app_search_goodsActivity.this, "网络连接良好", Toast.LENGTH_SHORT).show();
                    MyRetrofit.serviceAPI.searchProduct(query).enqueue(new Callback<SearchResult>() {
                        @Override
                        public void onResponse(Call<SearchResult> call, Response<SearchResult> response) {
                            if (response.isSuccessful()) {
                                response.body().getData().forEach(new Consumer<ProductSimpleInfoResult.CommodityInfo>() {
                                    @Override
                                    public void accept(ProductSimpleInfoResult.CommodityInfo commodityInfo) {
                                        list.add(new ProductSimple(commodityInfo.getCommodity_identity(), commodityInfo.getMedia().get(0).getImage(), commodityInfo.getTitle(), commodityInfo.getInformation(), String.valueOf(commodityInfo.getPrice()), String.valueOf(commodityInfo.getRating()), String.valueOf(commodityInfo.getID())));
                                    }
                                });
                                recyclerViewAdapter.notifyDataSetChanged();
                            } else {
                                Toast.makeText(app_search_goodsActivity.this, "有问题", Toast.LENGTH_SHORT).show();
                            }
                        }

                        @Override
                        public void onFailure(Call<SearchResult> call, Throwable t) {
                            Toast.makeText(app_search_goodsActivity.this, "失败", Toast.LENGTH_SHORT).show();
                        }
                    });

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