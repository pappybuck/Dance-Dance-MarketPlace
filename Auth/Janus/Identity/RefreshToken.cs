using System.ComponentModel.DataAnnotations;

namespace Janus.Identity
{
    public class RefreshToken
    {
        [Key]
        public string Token { get; set; }
        public string id { get; set; }
        public DateTime TokenExpire { get; set; }
        
        public DateTime TokenCreated { get; set; }
        public DateTime? Revoked { get; set; }
        public bool Invalidated { get; set; }
        public string UserId { get; set; }
        public User User { get; set; }
    }
}
