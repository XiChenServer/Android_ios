//
//  MineViewController.m
//  three
//
//  Created by Huc on 2023/12/10.
//

#import "MineViewController.h"
#import "DataViewController.h"
#import "LoginViewController.h"
#import "MineTableViewCell.h"
#import "SettingTableViewCell.h"

@interface MineViewController ()

@end

@implementation MineViewController

- (void)viewWillAppear:(BOOL)animated {
    if ([self.navigationController.tabBarController.title isEqualToString:@"KKK"]) {
        self.TOKEN = @"";
    } else {
        if (![self.TOKEN isEqualToString:self.navigationController.tabBarController.title]) {
            self.TOKEN = self.navigationController.tabBarController.title;
            [self getUserInformationWithTOKEN:self.TOKEN];
        }
        [self.tableView reloadData];
    }
}

- (void)loadTableView {
    self.tableView = [[UITableView alloc] initWithFrame:self.view.bounds style:UITableViewStyleGrouped];
    self.tableView.sectionFooterHeight = 4;
    self.tableView.sectionHeaderHeight = 4;
    self.tableView.delegate = self;
    self.tableView.dataSource = self;
    self.tableView.showsVerticalScrollIndicator = NO;
    [self.view addSubview: self.tableView];
    
    [self.tableView registerClass:[MineTableViewCell class] forCellReuseIdentifier:@"mineCell"];
    [self.tableView registerClass:[SettingTableViewCell class] forCellReuseIdentifier:@"setCell"];
    
//    UIImageView *imageView = [[UIImageView alloc]init];
//    imageView.frame = CGRectMake(0, 0, self.view.bounds.size.width, 300);
//    imageView.backgroundColor = UIColor.whiteColor;
//    self.tableView.backgroundView = [[UIImageView alloc]init];
//    [self.tableView.backgroundView addSubview:imageView];
    
    self.data = @[@[@"服务"], @[@"发布", @"购买", @"关注", @"收藏"], @[@"设置"]];
}

- (void)viewDidLoad {
    [super viewDidLoad];
    [self loadUserData];
    [self loadTableView];
    // Do any additional setup after loading the view.
}

- (void)loadUserData {
    if ([self.navigationController.tabBarController.title isEqualToString:@"KKK"]) {
        self.TOKEN = @"";
    } else {
        self.TOKEN = self.navigationController.tabBarController.title;
    }
}

- (UITableViewCell *)tableView:(UITableView *)tableView cellForRowAtIndexPath:(NSIndexPath *)indexPath {
    if (!indexPath.section) {
        MineTableViewCell *cell = [tableView dequeueReusableCellWithIdentifier:@"mineCell" forIndexPath:indexPath];
        cell.accessoryType = UITableViewCellAccessoryDisclosureIndicator;
        return cell;
    }
    
    SettingTableViewCell *cell = [tableView dequeueReusableCellWithIdentifier:@"setCell" forIndexPath:indexPath];
    
    cell.mainLabel.text = self.data[indexPath.section - 1][indexPath.row];
    
    cell.accessoryType = UITableViewCellAccessoryDisclosureIndicator;
    
    return cell;
}

- (void)tableView:(UITableView *)tableView didSelectRowAtIndexPath:(NSIndexPath *)indexPath {
    [tableView deselectRowAtIndexPath: indexPath animated: YES];
    if (indexPath.section != 3 && !self.TOKEN.length) {
        [self.navigationController pushViewController:[[LoginViewController alloc] init] animated:YES];
    } else if (!indexPath.section) {
        [self.navigationController pushViewController:[[DataViewController alloc] init] animated:YES];
    }
    
}

- (NSInteger)numberOfSectionsInTableView:(UITableView *)tableView {
    return self.data.count + 1;
}

- (NSInteger)tableView:(UITableView *)tableView numberOfRowsInSection:(NSInteger)section {
    if (section) {
        NSArray *temp = self.data[section - 1];
        return temp.count;
    } else {
        return 1;
    }
}

- (CGFloat)tableView:(UITableView *)tableView heightForRowAtIndexPath:(NSIndexPath *)indexPath {
    if (indexPath.section) {
        return 60;
    } else {
        return 160;
    }
}

- (void)getUserInformationWithTOKEN:(NSString *)token {
    NSURLComponents *urlComponents = [NSURLComponents componentsWithString:@"http://192.168.0.146:13000/user/get/info"];
    NSMutableArray *queryItems = [NSMutableArray array];

    // 创建要发送的JSON参数
    NSString *str = [NSString stringWithFormat:@"Bearer %@", token];
    NSURL *url = urlComponents.URL;
    NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:url];
    request.HTTPMethod = @"GET";
    [request setValue:str forHTTPHeaderField:@"Authorization"];
    NSURLSession *session = [NSURLSession sharedSession];
    NSURLSessionDataTask *dataTask = [session dataTaskWithRequest:request completionHandler:^(NSData *data, NSURLResponse *response, NSError *error) {
        if (error) {
            NSLog(@"Failed to send request: %@", error);
            return;
        }
        
        // 处理响应数据
        NSString *responseData = [[NSString alloc] initWithData:data encoding:NSUTF8StringEncoding];
        NSLog(@"Response: %@", responseData);
    }];

    [dataTask resume];
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
