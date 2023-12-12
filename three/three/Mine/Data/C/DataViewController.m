//
//  DataViewController.m
//  three
//
//  Created by Huc on 2023/12/10.
//

#import "DataViewController.h"
#import "AlterViewController.h"
#import <AVFoundation/AVFoundation.h>
#import <MobileCoreServices/MobileCoreServices.h>
#import <MediaPlayer/MediaPlayer.h>
#import <AVKit/AVKit.h>

@interface DataViewController ()

@end

@implementation DataViewController

- (void)loadTableView {
    
    self.tableView = [[UITableView alloc] initWithFrame:self.view.bounds style:UITableViewStyleGrouped];
    self.tableView.sectionFooterHeight = 4;
    self.tableView.sectionHeaderHeight = 4;
    self.tableView.delegate = self;
    self.tableView.dataSource = self;
    self.tableView.showsVerticalScrollIndicator = NO;
    [self.view addSubview: self.tableView];
    
    self.data = @[@[@"头像", @"昵称"], @[@"手机号", @"邮箱号", @"微信号"], @[@"我的地址", @"更改密码"]];
}

- (void)viewDidLoad {
    [super viewDidLoad];
    self.title = @"个人信息";
    [self loadTableView];
    // Do any additional setup after loading the view.
}

- (UITableViewCell *)tableView:(UITableView *)tableView cellForRowAtIndexPath:(NSIndexPath *)indexPath {
    
    UITableViewCell *cell = [[UITableViewCell alloc] init];
    
    cell.textLabel.text = self.data[indexPath.section][indexPath.row];
    [cell.contentView addSubview:cell.textLabel];
    
    cell.accessoryType = UITableViewCellAccessoryDisclosureIndicator;
    
    return cell;
}

- (void)tableView:(UITableView *)tableView didSelectRowAtIndexPath:(NSIndexPath *)indexPath {
    [tableView deselectRowAtIndexPath: indexPath animated: YES];
    if (indexPath.section == 0 && indexPath.row == 0) {
        UIImagePickerController *cameraPicker = [[UIImagePickerController alloc] init];
        cameraPicker.delegate = self;
        cameraPicker.allowsEditing = YES;
        cameraPicker.sourceType = UIImagePickerControllerSourceTypePhotoLibrary;
              //跳转到UIImagePickerController控制器弹出相册
        cameraPicker.modalPresentationStyle = UIModalPresentationFullScreen;
        cameraPicker.modalTransitionStyle = UIModalTransitionStyleFlipHorizontal;
        [self presentViewController:cameraPicker animated:YES completion:nil];
        return;
    }
    if (indexPath.section == 2 && indexPath.row == 0) {
        
        return;;
    }
    AlterViewController *alter = [[AlterViewController alloc] init];
    alter.row = indexPath.section*10 + indexPath.row;
    [self.navigationController pushViewController:alter animated:YES];
    
}

- (void)imagePickerController:(UIImagePickerController *)picker didFinishPickingMediaWithInfo:(NSDictionary *)info{
    [picker dismissViewControllerAnimated:YES completion:nil];
}

- (NSInteger)numberOfSectionsInTableView:(UITableView *)tableView {
    return self.data.count;
}

- (NSInteger)tableView:(UITableView *)tableView numberOfRowsInSection:(NSInteger)section {
    NSArray *temp = self.data[section];
    return temp.count;
}

- (CGFloat)tableView:(UITableView *)tableView heightForRowAtIndexPath:(NSIndexPath *)indexPath {
    return 60;
}

/*
#pragma mark - Navigation

// In a storyboard-based application, you will often want to do a little preparation before navigation
- (void)prepareForSegue:(UIStoryboardSegue *)segue sender:(id)sender {
    // Get the new view controller using [segue destinationViewController].
    // Pass the selected object to the new view controller.
}
*/

@end
