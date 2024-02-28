package com.example.mine_pager.fragment;

import android.os.Bundle;
import android.os.Handler;
import android.os.HandlerThread;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.MenuItem;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.recyclerview.widget.GridLayoutManager;

import com.alibaba.android.arouter.launcher.ARouter;
import com.example.common.room.AppDatabase;
import com.example.common.room.dao.ProductSimpleDao;
import com.example.common.room.entitues.Product;
import com.example.common.room.entitues.ProductSimple;
import com.example.common.room.entitues.ShoppingCarOrder;
import com.example.core_net_work.MyRetrofit;
import com.example.core_net_work.model.goods.ProductSimpleInfoResult;
import com.example.mine_pager.R;
import com.example.mine_pager.adapter.RecyclerViewAdapter;
import com.example.mine_pager.databinding.FragmentBlankBinding;
import com.github.jdsjlzx.interfaces.OnItemLongClickListener;
import com.github.jdsjlzx.interfaces.OnRefreshListener;
import com.github.jdsjlzx.recyclerview.LRecyclerViewAdapter;
import com.github.jdsjlzx.recyclerview.ProgressStyle;
import com.google.android.material.snackbar.Snackbar;
import com.kennyc.bottomsheet.BottomSheetListener;
import com.kennyc.bottomsheet.BottomSheetMenuDialogFragment;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;
import java.util.function.Consumer;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;


public class BlankFragment extends Fragment {
    FragmentBlankBinding binding;
    private int pageSize = 10;
    int count = 1;

    private void requestData() {

    }

    public List<ProductSimple> list = new ArrayList();


    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_blank, container, false);
        binding = FragmentBlankBinding.bind(view);
//        list.add(new ProductSimple(String.valueOf(com.example.common.R.drawable.loading), "Loading", "loading", "loading", "0"));
//        list.add(new ProductSimple(String.valueOf(com.example.common.R.drawable.loading), "Loading", "loading", "loading", "0"));
//        list.add(new ProductSimple(String.valueOf(com.example.common.R.drawable.loading), "Loading", "loading", "loading", "0"));
//        list.add(new ProductSimple(String.valueOf(com.example.common.R.drawable.loading), "Loading", "loading", "loading", "0"));
//         创建 Handler 对象，并将 Looper 对象传入
        RecyclerViewAdapter recyclerViewAdapter = new RecyclerViewAdapter(list, getActivity());
        LRecyclerViewAdapter adapter = new LRecyclerViewAdapter(recyclerViewAdapter);
        GridLayoutManager gridLayoutManager = new GridLayoutManager(getActivity(), 2);
        binding.sellPagerRecyclerView.setLayoutManager(gridLayoutManager);

        binding.sellPagerRecyclerView.setAdapter(adapter);
        binding.sellPagerRecyclerView.refreshComplete(pageSize);
        //setEmptyView会导致
//        binding.sellPagerRecyclerView.setEmptyView(binding.emptyView.getRoot());

//        binding.sellPagerRecyclerView.setOnNetWorkErrorListener(new OnNetWorkErrorListener() {
//            @Override
//            public void reload() {
//
//            }
//        });

//        HandlerThread handlerThread = new HandlerThread("MyHandlerThread");
//        handlerThread.start();
//        Handler handler = new Handler(handlerThread.getLooper());
//        handler.post(new Runnable() {
//            @Override
//            public void run() {
//                List<ProductSimple> allInfo = AppDatabase.getInstance(getActivity()).productSimpleDao().getAllInfo();
//                list.clear();
//                list.addAll(allInfo);
//                adapter.notifyDataSetChanged();
//                binding.sellPagerRecyclerView.refreshComplete(pageSize);
//            }
//        });
        binding.sellPagerRecyclerView.setOnRefreshListener(new OnRefreshListener() {
            @Override
            public void onRefresh() {
                MyRetrofit.serviceAPI.getProductSimpleInfo().enqueue(new Callback<ProductSimpleInfoResult>() {
                    @Override
                    public void onResponse(Call<ProductSimpleInfoResult> call, Response<ProductSimpleInfoResult> response) {
                        if (response.isSuccessful()) {
                            Toast.makeText(getActivity(), "成功了", Toast.LENGTH_SHORT).show();
                            response.body().getData().getCommodity().forEach(new Consumer<ProductSimpleInfoResult.CommodityInfo>() {
                                @Override
                                public void accept(ProductSimpleInfoResult.CommodityInfo commodityInfo) {
//                                    list.clear();
                                    Log.d("大问题", commodityInfo.toString());
                                    List<ProductSimpleInfoResult.CommodityInfo.MediaResult> media = commodityInfo.getMedia();
//                                    if(){

//                                    }
//                                    list.clear();
                                    list.add(new ProductSimple(media.get(0).getImage(), commodityInfo.getTitle(), commodityInfo.getInformation(), String.valueOf(commodityInfo.getPrice()), "3.5"));
                                    binding.sellPagerRecyclerView.refreshComplete(pageSize);
                                    adapter.notifyDataSetChanged();
                                }
                            });
                            HandlerThread handlerThread = new HandlerThread("saveProductSimple");
                            handlerThread.start();
                            Handler handler = new Handler(handlerThread.getLooper());
                            handler.post(new Runnable() {
                                @Override
                                public void run() {
                                    ProductSimpleDao simpleDao = AppDatabase.getInstance(getActivity()).productSimpleDao();
                                    simpleDao.deleteAll();
                                    simpleDao.insertAll(list);
                                }
                            });
                        } else {
                            Toast.makeText(getActivity(), "有问题", Toast.LENGTH_SHORT).show();
                        }
                        binding.sellPagerRecyclerView.refreshComplete(pageSize);
                    }

                    @Override
                    public void onFailure(Call<ProductSimpleInfoResult> call, Throwable t) {
                        binding.sellPagerRecyclerView.refreshComplete(pageSize);
                        Toast.makeText(getActivity(), "失败了", Toast.LENGTH_SHORT).show();
                        Log.d("大问题", t.toString());
                    }
                });
            }
        });

        //强制刷新
        binding.sellPagerRecyclerView.setRefreshProgressStyle(ProgressStyle.Pacman);
        binding.sellPagerRecyclerView.forceToRefresh();
        binding.sellPagerRecyclerView.refreshComplete(pageSize);
        adapter.notifyDataSetChanged();

//        binding.sellPagerRecyclerView.setLoadMoreEnabled(true);
//        binding.sellPagerRecyclerView.setLoadMoreFooter(new LoadingFooter(getActivity()), true);
//        binding.sellPagerRecyclerView.setFooterViewHint("拼命加载中", "已经全部为你呈现了", "网络不给力啊，点击再试一次吧");
//        binding.sellPagerRecyclerView.setOnLoadMoreListener(new OnLoadMoreListener() {
//            @Override
//            public void onLoadMore() {
//                Toast.makeText(getActivity(), "wertwert", Toast.LENGTH_SHORT).show();
//                Log.d("loadmore问题", String.valueOf(count));
//                binding.sellPagerRecyclerView.forceToRefresh();
//                count++;
//
//            }
//        });

        adapter.setOnItemLongClickListener(new OnItemLongClickListener() {
            @Override
            public void onItemLongClick(View view, int position) {
                new BottomSheetMenuDialogFragment.Builder(getActivity()).setSheet(R.menu.menu_more).setTitle("more").setListener(new BottomSheetListener() {
                    @Override
                    public void onSheetShown(@NonNull BottomSheetMenuDialogFragment bottomSheetMenuDialogFragment, @Nullable Object o) {

                    }

                    @Override
                    public void onSheetItemSelected(@NonNull BottomSheetMenuDialogFragment bottomSheetMenuDialogFragment, @NonNull MenuItem menuItem, @Nullable Object o) {
                        if (menuItem.getItemId() == R.id.more_shopping_car) {
                            HandlerThread shoppingCar = new HandlerThread("readToShoppingCar");
                            shoppingCar.start();
                            Handler thread = new Handler(shoppingCar.getLooper());
                            thread.post(new Runnable() {
                                @Override
                                public void run() {
                                    List<ShoppingCarOrder> temp = new ArrayList<>();
                                    if (temp.size() > 0) {
                                        String productIdentity = list.get(position).getProduct_identity();
//                                    //通过唯一标识，找到这个商品，具体在哪里
                                        try {
                                            List<Product> product = AppDatabase.getInstance(getActivity()).productDao().getProduct(productIdentity);
                                            Product product1 = product.get(0);
                                            temp.add(new ShoppingCarOrder(product1.getImage(), Integer.valueOf(product1.getPrice()), 1, product1.getTitle()));
                                            AppDatabase.getInstance(getActivity()).shoppingCarDao().insertAll(temp);
                                        } catch (NumberFormatException e) {

                                            Log.d("问题", e.toString());
                                        }

                                        getActivity().runOnUiThread(new Runnable() {
                                            @Override
                                            public void run() {
                                                Snackbar.make(view, "添加成功", Snackbar.LENGTH_LONG).show();
                                            }
                                        });
                                    } else {
                                        Snackbar.make(view, "无商品可添加", Snackbar.LENGTH_LONG).show();
                                    }
                                }
                            });
                        } else if (menuItem.getItemId() == R.id.more_collect) {

                        }
                    }

                    @Override
                    public void onSheetDismissed(@NonNull BottomSheetMenuDialogFragment bottomSheetMenuDialogFragment, @Nullable Object o, int i) {

                    }
                }).show(getActivity().getSupportFragmentManager());

            }
        });
        binding.imageButtonSort.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new BottomSheetMenuDialogFragment.Builder(getActivity()).setSheet(R.menu.menu_sort).setTitle("Sort").setListener(new BottomSheetListener() {
                    @Override
                    public void onSheetShown(@NonNull BottomSheetMenuDialogFragment bottomSheetMenuDialogFragment, @Nullable Object o) {

                    }

                    @Override
                    public void onSheetItemSelected(@NonNull BottomSheetMenuDialogFragment bottomSheetMenuDialogFragment, @NonNull MenuItem menuItem, @Nullable Object o) {
                        if (menuItem.getItemId() == R.id.sort_max_min) {
                            //从高到低
                            list.sort(new Comparator<ProductSimple>() {
                                @Override
                                public int compare(ProductSimple o1, ProductSimple o2) {
                                    if (Float.valueOf(o1.getPrice()) - Float.valueOf(o2.getPrice()) < 0) {
                                        return 1;
                                    }
                                    return -1;
                                }
                            });
                            recyclerViewAdapter.notifyDataSetChanged();
                        } else if (menuItem.getItemId() == R.id.sort_min_max) {
                            //从低到高
                            list.sort(new Comparator<ProductSimple>() {
                                @Override
                                public int compare(ProductSimple o1, ProductSimple o2) {
                                    if (Float.valueOf(o1.getPrice()) - Float.valueOf(o2.getPrice()) > 0) {
                                        return 1;
                                    }
                                    return -1;
                                }
                            });
                            recyclerViewAdapter.notifyDataSetChanged();
                        }
                    }

                    @Override
                    public void onSheetDismissed(@NonNull BottomSheetMenuDialogFragment bottomSheetMenuDialogFragment, @Nullable Object o, int i) {

                    }
                }).show(getActivity().getSupportFragmentManager());
            }
        });
        binding.imageButtonFilter.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                //直接新建一个碎片
            }
        });
        binding.sellPagerFramelayout.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                //直接打开活动
//                之前使用的碎片的操作
//                FragmentManager fragmentManager = getActivity().getSupportFragmentManager();
//                FragmentTransaction fragmentTransaction = fragmentManager.beginTransaction();
//                SearchFragment fragment = new SearchFragment();
//                fragmentTransaction.replace(, fragment); fragmentTransaction.addToBackStack(null);
//                fragmentTransaction.commit();
                ARouter.getInstance().build("/mine_pager/app_search_goodsActivity").navigation();
                getActivity().overridePendingTransition(com.example.common.R.anim.set_in, com.example.common.R.anim.set_out);
            }
        });
        return view;
    }
}