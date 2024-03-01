package com.example.mine.fragment.listviewfragment;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.viewpager2.widget.ViewPager2;

import com.example.mine.R;
import com.example.mine.databinding.FragmentProductBinding;
import com.example.mine.fragment.listviewfragment.adapter.ViewPager2Adapter;
import com.example.mine.fragment.listviewfragment.productfragment.BoughtFragment;
import com.example.mine.fragment.listviewfragment.productfragment.ProductSellIngFragment;
import com.example.mine.fragment.listviewfragment.productfragment.SellOutOrderFragment;
import com.google.android.material.tabs.TabLayout;

import java.util.ArrayList;
import java.util.List;

/**
 * @Author winiymissl
 * @Date 2024-02-23 16:52
 * @Version 1.0
 */
public class ProductFragment extends Fragment {
    FragmentProductBinding binding;

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_product, container, false);

        binding = FragmentProductBinding.bind(view);
        binding.transparentViewProductProduct.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {

            }
        });

        List<Fragment> list = new ArrayList<>();
        list.add(new ProductSellIngFragment("已上架"));
        list.add(new BoughtFragment("购买订单"));
        list.add(new SellOutOrderFragment("出售记录"));
        //按顺序添加tab，按照设计的顺序
        TabLayout.Tab tab_up = binding.productTablayoutProduct.newTab();
        TabLayout.Tab tab_bought = binding.productTablayoutProduct.newTab();
        TabLayout.Tab tab_sell_order = binding.productTablayoutProduct.newTab();
        tab_up.setText("已上架");
        tab_bought.setText("购买记录");
        tab_sell_order.setText("出售记录");
        binding.productTablayoutProduct.addTab(tab_up);
        binding.productTablayoutProduct.addTab(tab_bought);
        binding.productTablayoutProduct.addTab(tab_sell_order);
        ViewPager2Adapter adapter = new ViewPager2Adapter(getChildFragmentManager(), getLifecycle(), list);
        binding.productViewpagerProduct.setAdapter(adapter);
        binding.productViewpagerProduct.registerOnPageChangeCallback(new ViewPager2.OnPageChangeCallback() {
            @Override
            public void onPageSelected(int position) {
                super.onPageSelected(position);
                if (position == 0) {
                    binding.productTablayoutProduct.selectTab(tab_up);
                } else if (position == 1) {
                    binding.productTablayoutProduct.selectTab(tab_bought);
                } else if (position == 2) {
                    binding.productTablayoutProduct.selectTab(tab_sell_order);
                }
            }
        });
        binding.productViewpagerProduct.setOnScrollChangeListener(new View.OnScrollChangeListener() {
            
            @Override
            public void onScrollChange(View v, int scrollX, int scrollY, int oldScrollX, int oldScrollY) {

            }
        });
        binding.productTablayoutProduct.addOnTabSelectedListener(new TabLayout.OnTabSelectedListener() {
            @Override
            public void onTabSelected(TabLayout.Tab tab) {
                binding.productViewpagerProduct.setCurrentItem(tab.getPosition(), true);
            }

            @Override
            public void onTabUnselected(TabLayout.Tab tab) {

            }

            @Override
            public void onTabReselected(TabLayout.Tab tab) {

            }
        });

//        for (int i = 0; i < binding.tabLayoutProduct.getTabCount(); i++) {
//            TabLayout.Tab tabAt = binding.tabLayoutProduct.getTabAt(i);
//            if (fragmentList.get(i) instanceof ProductSellIngFragment) {
//                tabAt.setText("已上架");
//            }
//            if (fragmentList.get(i) instanceof HistoryOrderFragment) {
//                tabAt.setText("历史订单");
//            }
//        }
//        List<ProductEntity> list = new ArrayList<>();
//        ProductRecyclerViewAdapter adapter = new ProductRecyclerViewAdapter(list, getActivity());
//        adapterLR = new LRecyclerViewAdapter(adapter);
//        binding.lrecyclerviewProduct.setLayoutManager(new GridLayoutManager(getActivity(), 1));
//        binding.lrecyclerviewProduct.setAdapter(adapterLR);
//        HandlerThread handlerThread = new HandlerThread("DatabaseThread");
//        handlerThread.start();
//        Handler handler = new Handler(handlerThread.getLooper());
//        handler.post(new Runnable() {
//            @Override
//            public void run() {
//                List<User> allInfo = AppDatabase.getInstance(getActivity()).userDao().getAllInfo();
//                User user = allInfo.get(0);
//                MyRetrofit.serviceAPI.getUserUploaded(user.getUser_identity()).enqueue(new Callback<UserLoadedResult>() {
//                    @Override
//                    public void onResponse(Call<UserLoadedResult> call, Response<UserLoadedResult> response) {
//                        if (response.isSuccessful()) {
//                            UserLoadedResult.UserLoadedResultData data = response.body().getData();
//                            List<ProductSimpleInfoResult.CommodityInfo> commodity = data.getUser_products().getCommodity();
//                            if (commodity != null) {
////                                try {
//                                commodity.forEach(new Consumer<ProductSimpleInfoResult.CommodityInfo>() {
//                                    @Override
//                                    public void accept(ProductSimpleInfoResult.CommodityInfo commodityInfo) {
//                                        list.add(new ProductEntity(commodityInfo.getMedia().get(0).getImage(), commodityInfo.getTitle(), String.valueOf(commodityInfo.getPrice())));
//                                    }
//                                });
////                                } catch (Exception e) {
////                                    Log.d("恶心问题",e.toString());
////                                }
//                                adapter.notifyDataSetChanged();
//                                binding.progressBarProduct.setVisibility(View.GONE);
//                                binding.lrecyclerviewProduct.setVisibility(View.VISIBLE);
//                            } else {
//                                binding.progressBarProduct.setVisibility(View.GONE);
//                                binding.textViewSelling.setVisibility(View.VISIBLE);
//                            }
//                        }
//                    }
//
//                    @Override
//                    public void onFailure(Call<UserLoadedResult> call, Throwable t) {
//                        Log.d("问题", t.toString());
//                        Toast.makeText(getActivity(), "失败", Toast.LENGTH_SHORT).show();
//                    }
//                });
//            }
//        });
//
//        binding.lrecyclerviewProduct.setOnRefreshListener(new OnRefreshListener() {
//            @Override
//            public void onRefresh() {
//                HandlerThread handlerThread = new HandlerThread("DatabaseThread");
//                handlerThread.start();
//                Handler handler = new Handler(handlerThread.getLooper());
//                handler.post(new Runnable() {
//                    @Override
//                    public void run() {
//                        List<User> allInfo = AppDatabase.getInstance(getActivity()).userDao().getAllInfo();
//                        User user = allInfo.get(0);
//                        MyRetrofit.serviceAPI.getUserUploaded(user.getUser_identity()).enqueue(new Callback<UserLoadedResult>() {
//                            @Override
//                            public void onResponse(Call<UserLoadedResult> call, Response<UserLoadedResult> response) {
//                                if (response.isSuccessful()) {
//                                    UserLoadedResult.UserLoadedResultData data = response.body().getData();
//                                    List<ProductSimpleInfoResult.CommodityInfo> commodity = data.getUser_products().getCommodity();
//                                    if (commodity != null) {
////                                try {
//                                        commodity.forEach(new Consumer<ProductSimpleInfoResult.CommodityInfo>() {
//                                            @Override
//                                            public void accept(ProductSimpleInfoResult.CommodityInfo commodityInfo) {
//                                                list.clear();
//                                                list.add(new ProductEntity(commodityInfo.getMedia().get(0).getImage(), commodityInfo.getTitle(), String.valueOf(commodityInfo.getPrice())));
//                                            }
//                                        });
////                                } catch (Exception e) {
////                                    Log.d("恶心问题",e.toString());
////                                }
//                                        adapterLR.notifyDataSetChanged();
//                                    } else {
//                                        binding.textViewSelling.setVisibility(View.VISIBLE);
//                                        binding.lrecyclerviewProduct.setVisibility(View.GONE);
//                                    }
//                                    binding.lrecyclerviewProduct.refreshComplete(10);
//                                }
//                            }
//
//                            @Override
//                            public void onFailure(Call<UserLoadedResult> call, Throwable t) {
//
//                            }
//                        });
//                    }
//                });
//            }
//        });
//
//        binding.floatingActionButton.setOnClickListener(new View.OnClickListener() {
//            @Override
//            public void onClick(View v) {
//                FragmentManager fragmentManager = getActivity().getSupportFragmentManager();
//                FragmentTransaction fragmentTransaction = fragmentManager.beginTransaction();
//                AddProductFragment fragment = new AddProductFragment();
//                fragmentTransaction.setCustomAnimations(com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right, com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right);
//                fragmentTransaction.add(R.id.frame_mine, fragment);
//                fragmentTransaction.addToBackStack(null);
//                fragmentTransaction.commit();
//            }
//        });
        return view;
    }
}
