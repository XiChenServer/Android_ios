package com.example.mine.fragment;

import static android.app.Activity.RESULT_OK;

import android.Manifest;
import android.content.Intent;
import android.graphics.Color;
import android.net.Uri;
import android.os.Bundle;
import android.os.Handler;
import android.os.HandlerThread;
import android.provider.MediaStore;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.MenuItem;
import android.view.View;
import android.view.ViewGroup;
import android.widget.AdapterView;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentManager;
import androidx.fragment.app.FragmentTransaction;
import androidx.swiperefreshlayout.widget.SwipeRefreshLayout;

import com.alibaba.android.arouter.facade.annotation.Route;
import com.alibaba.android.arouter.launcher.ARouter;
import com.bumptech.glide.Glide;
import com.example.common.room.AppDatabase;
import com.example.common.room.dao.UserDao;
import com.example.common.room.entitues.User;
import com.example.core_net_work.MyRetrofit;
import com.example.core_net_work.model.userInfo.UserInfoResult;
import com.example.mine.R;
import com.example.mine.adpater.MineListViewAdapter;
import com.example.mine.adpater.impl.MineListViewAdapterImpl;
import com.example.mine.databinding.AppMineFragmentBinding;
import com.example.mine.fragment.listviewfragment.ProductFragment;
import com.example.mine.fragment.listviewfragment.SettingFragment;
import com.google.android.material.snackbar.Snackbar;
import com.kennyc.bottomsheet.BottomSheetListener;
import com.kennyc.bottomsheet.BottomSheetMenuDialogFragment;
import com.tencent.mmkv.MMKV;

import java.util.List;

import pub.devrel.easypermissions.EasyPermissions;
import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

/**
 * @Author winiymissl
 * @Date 2023-12-12 17:54
 * @Version 1.0
 */
@Route(path = "/mine/fragment/app_mine_MainFragment")
public class app_mine_MainFragment extends Fragment {
    AppMineFragmentBinding binding;
    private static final int REQUEST_CODE_SELECT_IMAGE = 1;
    private static final int REQUEST_CODE_PERMISSIONS = 2;
    private String[] permissions = {Manifest.permission.READ_EXTERNAL_STORAGE};

    public void openGallery() {
        Intent intent = new Intent(Intent.ACTION_PICK, MediaStore.Images.Media.EXTERNAL_CONTENT_URI);
        startActivityForResult(intent, REQUEST_CODE_SELECT_IMAGE);
    }

    private void requestPermission() {
        if (EasyPermissions.hasPermissions(getActivity(), permissions)) {
            // 如果权限已授予，执行您的操作
            // 比如打开相机等

            Log.d("ThereIsProblem", "已通过申请");
        } else {
            // 如果权限未授予，请求权限
            EasyPermissions.requestPermissions(getActivity(), "提供相机权限", REQUEST_CODE_PERMISSIONS, permissions);
        }
    }

    @Override
    public void onActivityResult(int requestCode, int resultCode, @Nullable Intent data) {
        super.onActivityResult(requestCode, resultCode, data);
        if (requestCode == REQUEST_CODE_SELECT_IMAGE && resultCode == RESULT_OK && data != null) {
            Uri uri = data.getData();
            Glide.with(getActivity()).load(uri).into(binding.ivMine);
            //将这个uri存储在mmkv中
            MMKV.defaultMMKV().encode("imageBackground", uri.toString());
            //将字体修改
            binding.textViewName.setTextColor(Color.WHITE);
            binding.textViewPhone.setTextColor(Color.WHITE);
        }
    }

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {

        View view = inflater.inflate(R.layout.app_mine_fragment, container, false);
        binding = AppMineFragmentBinding.bind(view);
        MineListViewAdapter adapter = new MineListViewAdapter(MineListViewAdapterImpl.getList());
        binding.mineListview.setAdapter(adapter);
        MMKV mmkv = MMKV.defaultMMKV();
        requestPermission();
        //根据mmkv是否存储，进行修改
        if (mmkv.contains("imageBackground")) {
            //加载背景照片
            Glide.with(getActivity()).load(mmkv.decodeString("imageBackground")).error(com.example.common.R.drawable.climb).into(binding.ivMine);
            //根据有无背景修改字体的颜色
            binding.textViewName.setTextColor(Color.WHITE);
            binding.textViewPhone.setTextColor(Color.WHITE);
        }

        binding.refreshUserInfo.setOnRefreshListener(new SwipeRefreshLayout.OnRefreshListener() {
            @Override
            public void onRefresh() {
                MyRetrofit.serviceAPI.getUserInfo("bearer " + mmkv.getString("token", null)).enqueue(new Callback<UserInfoResult>() {
                    @Override
                    public void onResponse(Call<UserInfoResult> call, Response<UserInfoResult> response) {
                        if (response.isSuccessful()) {
                            Toast.makeText(getActivity(), "刷新成功", Toast.LENGTH_SHORT).show();
                            requestData(response.body().getData());
                        } else {
                            Toast.makeText(getActivity(), "刷新有问题", Toast.LENGTH_SHORT).show();
                        }
                        binding.refreshUserInfo.setRefreshing(false);
                    }

                    @Override
                    public void onFailure(Call<UserInfoResult> call, Throwable t) {
                        Toast.makeText(getActivity(), "刷新失败", Toast.LENGTH_SHORT).show();
                        binding.refreshUserInfo.setRefreshing(false);
                    }
                });
            }
        });
        binding.mineListview.setOnItemClickListener(new AdapterView.OnItemClickListener() {
            @Override
            public void onItemClick(AdapterView<?> parent, View view, int position, long id) {
                if (position == 0) {
                    Toast.makeText(getActivity(), "商品", Toast.LENGTH_SHORT).show();
                    FragmentManager fragmentManager = getActivity().getSupportFragmentManager();
                    FragmentTransaction fragmentTransaction = fragmentManager.beginTransaction();
                    ProductFragment fragment = new ProductFragment();
                    fragmentTransaction.setCustomAnimations(com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right, com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right);
                    fragmentTransaction.add(R.id.frame_mine, fragment);
                    fragmentTransaction.addToBackStack(null);
                    fragmentTransaction.commit();

                } else if (position == 1) {
                    try {
                        Toast.makeText(getActivity(), "设置", Toast.LENGTH_SHORT).show();
                        FragmentManager fragmentManager = getActivity().getSupportFragmentManager();
                        FragmentTransaction fragmentTransaction = fragmentManager.beginTransaction();
                        SettingFragment fragment = new SettingFragment();
                        fragmentTransaction.setCustomAnimations(com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right, com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right);
                        fragmentTransaction.replace(R.id.frame_mine, fragment);
                        fragmentTransaction.addToBackStack(null);
                        fragmentTransaction.commit();
                    } catch (Exception e) {
                        Log.d("这里有一个问题", e.toString());
                    }
                }
//                else if (id == 1) {
//                    Toast.makeText(getActivity(), "收藏", Toast.LENGTH_SHORT).show();
//                    FragmentManager fragmentManager = getActivity().getSupportFragmentManager();
//                    FragmentTransaction fragmentTransaction = fragmentManager.beginTransaction();
//                    CollectionFragment fragment = new CollectionFragment();
//                    fragmentTransaction.setCustomAnimations(com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right, com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right);
//                    fragmentTransaction.replace(R.id.frame_mine, fragment);
//                    fragmentTransaction.addToBackStack(null);
//                    fragmentTransaction.commit();
//                }
            }
        });
        binding.ivMine.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                //替换背景照片，或者
                new BottomSheetMenuDialogFragment.Builder(getActivity()).setSheet(R.menu.menu_bottom_sheet_menu).setTitle("More").setListener(new BottomSheetListener() {
                    @Override
                    public void onSheetShown(@NonNull BottomSheetMenuDialogFragment bottomSheetMenuDialogFragment, @Nullable Object o) {

                    }

                    @Override
                    public void onSheetItemSelected(@NonNull BottomSheetMenuDialogFragment bottomSheetMenuDialogFragment, @NonNull MenuItem menuItem, @Nullable Object o) {
                        if (menuItem.getItemId() == R.id.upload) {
                            openGallery();
                        } else if (menuItem.getItemId() == R.id.info_personal) {
                            FragmentManager fragmentManager = getActivity().getSupportFragmentManager();
                            FragmentTransaction fragmentTransaction = fragmentManager.beginTransaction();
                            app_mine_infoFragment fragment = new app_mine_infoFragment();
                            fragmentTransaction.setCustomAnimations(com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right, com.example.common.R.anim.slide_in_from_right, com.example.common.R.anim.slide_out_to_right);
                            fragmentTransaction.replace(R.id.frame_mine, fragment);
                            fragmentTransaction.addToBackStack(null);
                            fragmentTransaction.commit();
//                            ARouter.getInstance().build("/mine/fragment/app_mine_infoFragment").navigation();
                        }
                    }

                    @Override
                    public void onSheetDismissed(@NonNull BottomSheetMenuDialogFragment bottomSheetMenuDialogFragment, @Nullable Object o, int i) {

                    }
                }).show(getActivity().getSupportFragmentManager());
            }
        });
        binding.btnLogout.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                MMKV mmkv = MMKV.defaultMMKV();
                if (mmkv.getBoolean("token", false)) {
                    mmkv.remove("token");
                    binding.constraintIusLogin.setVisibility(View.GONE);
                    binding.textViewIsLogin.setVisibility(View.VISIBLE);
                }
                Snackbar.make(requireView(), "是否登录", 1000).setAction("去登陆", new View.OnClickListener() {
                    @Override
                    public void onClick(View v) {
                        ARouter.getInstance().build("/app_login/AppLoginMainActivity").withFlags(Intent.FLAG_ACTIVITY_NEW_TASK | Intent.FLAG_ACTIVITY_CLEAR_TASK).navigation();
                        getActivity().overridePendingTransition(com.example.common.R.anim.set_in, com.example.common.R.anim.set_out);
                    }
                }).show();
            }
        });
        //加载基本信息
        String token = MMKV.defaultMMKV().getString("token", null);
        if (token != null) {
            MyRetrofit.serviceAPI.getUserInfo("bearer " + token).enqueue(new Callback<UserInfoResult>() {
                @Override
                public void onResponse(Call<UserInfoResult> call, Response<UserInfoResult> response) {
                    if (response.isSuccessful()) {
                        UserInfoResult.UserInfoAllResult data = response.body().getData();
//                        Toast.makeText(getActivity(), "返回成功", Toast.LENGTH_SHORT).show();
                        Log.d("ThereIsProblem", "返回成功");
                        //得到从服务器返回的数据，保存到数据库中
//                        Log.d("ThereIsProblem", response.body().getData().getAvatar());
//                        Log.d("ThereIsProblem", response.body().getData().getName());
//                        Log.d("ThereIsProblem", response.body().getData().getNickname());
//                        Log.d("ThereIsProblem", response.body().getData().getAccount());
//                        Log.d("ThereIsProblem", response.body().getData().getPhone_number());
                        requestData(data);
                    } else {
                        Toast.makeText(getActivity(), "返回失败", Toast.LENGTH_SHORT).show();
                    }
                }

                @Override
                public void onFailure(Call<UserInfoResult> call, Throwable t) {
                    Toast.makeText(getActivity(), "返回失败" + t.toString(), Toast.LENGTH_SHORT).show();
                    Log.d("ThereIsProblem", t.toString());
                }
            });
        } else {
            Toast.makeText(getActivity(), "token为空", Toast.LENGTH_SHORT).show();
        }
        return view;
    }


    private void uploadInfo() {
        List<User> allInfo = AppDatabase.getInstance(getActivity()).userDao().getAllInfo();
        User user = allInfo.get(0);
        getActivity().runOnUiThread(new Runnable() {
            @Override
            public void run() {
                //加载头像
                Glide.with(getActivity()).load(user.getAvatar()).placeholder(com.example.common.R.drawable.loading).error(com.example.common.R.drawable.avatatloadfail).into(binding.imageView);
                //加载名字
                binding.textViewName.setText(user.getNickname());
                //加载电话号码
                binding.textViewPhone.setText(user.getPhone_number());
            }
        });
    }

    private void requestData(UserInfoResult.UserInfoAllResult data) {
        HandlerThread handlerThread = new HandlerThread("DatabaseThread");
        handlerThread.start();
        Handler handler = new Handler(handlerThread.getLooper());
        handler.post(new Runnable() {
            @Override
            public void run() {
                UserDao userDao = AppDatabase.getInstance(getActivity()).userDao();
                List<User> allInfo = userDao.getAllInfo();
                //只要有，就删掉，因为只有一个用户
                if (allInfo.size() != 0) {
                    userDao.delete(allInfo);
                }
                AppDatabase.getInstance(getActivity()).userDao().insertAll(new User(
                        data.getNickname(),
                        data.getAvatar(),
                        data.getEmail(),
                        data.getPhone_number(),
                        data.getName(),
                        data.getWechat_number(),
                        data.getUser_identity(),
                        data.getAccount(),
                        data.getAddress()
                ));                                //上传个人信息
                uploadInfo();
                // 注意：不要在这里更新 UI，因为这是在后台线程上执行的操作
            }
        });
    }
}
